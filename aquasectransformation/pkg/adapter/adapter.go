/*
Copyright 2022 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package aquasectransformation implements a CloudEvents adapter that...
package aquasectransformation

import (
	"context"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	targetce "github.com/triggermesh/triggermesh/pkg/targets/adapter/cloudevents"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type Event struct {
	ID          int    `json:"id"`
	Time        int    `json:"time"`
	Date        int    `json:"date"`
	Type        string `json:"type"`
	User        string `json:"user"`
	Action      string `json:"action"`
	Image       string `json:"image"`
	Imagehash   string `json:"imagehash"`
	Imageid     string `json:"imageid"`
	Container   string `json:"container"`
	Containerid string `json:"containerid"`
	Host        string `json:"host"`
	Hostid      string `json:"hostid"`
	Category    string `json:"category"`
	Result      int    `json:"result"`
	Data        string `json:"data"`
	AccountID   int    `json:"account_id"`
}

type CSNFEvent struct {
	Event struct {
		GUID             string    `json:"guid"`
		Name             string    `json:"name"`
		Severity         string    `json:"severity"`
		ShortDescription string    `json:"shortDescription"`
		StartTime        time.Time `json:"startTime"`
		Status           string    `json:"status"`
	} `json:"event"`
	Provider struct {
		AccountID string `json:"accountId"`
	} `json:"provider"`
	ProviderID   string `json:"providerId"`
	ProviderType string `json:"providerType"`
	Resource     struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Region     string `json:"region"`
		Type       string `json:"type"`
		Zone       string `json:"zone"`
	} `json:"resource"`
	Source struct {
		SourceID   string `json:"sourceId"`
		SourceName string `json:"sourceName"`
	} `json:"source"`
}

type envAccessor struct {
	pkgadapter.EnvConfig
	// BridgeIdentifier is the name of the bridge workflow this target is part of
	BridgeIdentifier string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
	// CloudEvents responses parametrization
	CloudEventPayloadPolicy string `envconfig:"EVENTS_PAYLOAD_POLICY" default:"error"`
	// Sink defines the target sink for the events. If no Sink is defined the
	// events are replied back to the sender.
	Sink string `envconfig:"K_SINK"`
}

// NewAdapter adapter implementation
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	env := envAcc.(*envAccessor)
	logger := logging.FromContext(ctx)

	replier, err := targetce.New(env.Component, logger.Named("replier"),
		targetce.ReplierWithStatefulHeaders(env.BridgeIdentifier),
		targetce.ReplierWithStaticResponseType("io.triggermesh.aquasectransformation.error"),
		targetce.ReplierWithPayloadPolicy(targetce.PayloadPolicy(env.CloudEventPayloadPolicy)))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	return &aquasectransformationadapter{
		sink:     env.Sink,
		replier:  replier,
		ceClient: ceClient,
		logger:   logger,
	}
}

var _ pkgadapter.Adapter = (*aquasectransformationadapter)(nil)

type aquasectransformationadapter struct {
	sink     string
	replier  *targetce.Replier
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger
}

// Start is a blocking function and will return if an error occurs
// or the context is cancelled.
func (a *aquasectransformationadapter) Start(ctx context.Context) error {
	a.logger.Info("Starting aquasectransformation Adapter")
	return a.ceClient.StartReceiver(ctx, a.dispatch)
}

func (a *aquasectransformationadapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	// a.logger.Infof("Received event: %v", event)

	ade := &Event{}
	if err := event.DataAs(&ade); err != nil {
		a.logger.Errorf("Error unmarshalling event: %v", err)
		return nil, nil
	}

	csnfEvent := &CSNFEvent{
		Event: struct {
			GUID             string    `json:"guid"`
			Name             string    `json:"name"`
			Severity         string    `json:"severity"`
			ShortDescription string    `json:"shortDescription"`
			StartTime        time.Time `json:"startTime"`
			Status           string    `json:"status"`
		}{
			GUID:             ade.Hostid,
			Name:             ade.Type,
			Severity:         "High",
			ShortDescription: "an aquasec short description goes here",
			StartTime:        time.Now(),
			Status:           "open",
		},
		Provider: struct {
			AccountID string `json:"accountId"`
		}{
			AccountID: string(ade.ID),
		},
		ProviderID:   string(ade.ID),
		ProviderType: "Aquasec",
		Resource: struct {
			Identifier string `json:"identifier"`
			Name       string `json:"name"`
			Region     string `json:"region"`
			Type       string `json:"type"`
			Zone       string `json:"zone"`
		}{
			Identifier: ade.Image,
			Name:       ade.Hostid,
			Region:     "us-1-c",
			Type:       "container",
			Zone:       "US",
		},
		Source: struct {
			SourceID   string `json:"sourceId"`
			SourceName string `json:"sourceName"`
		}{
			SourceID:   "Aquasec",
			SourceName: "Aquasec",
		},
	}

	csnfCEvent := cloudevents.NewEvent()
	csnfCEvent.SetID(csnfEvent.Event.GUID)
	csnfCEvent.SetType("io.triggermesh.csnf.oracle.aquasec.event")
	csnfCEvent.SetSource("Aquasec")
	csnfCEvent.SetSubject("Aquasec")
	csnfCEvent.SetDataContentType("application/json")

	if err := csnfCEvent.SetData(cloudevents.ApplicationJSON, csnfEvent); err != nil {
		a.logger.Errorf("Error setting event data: %v", err)
		return nil, nil
	}

	if a.sink != "" {
		a.logger.Infof("Sending event to sink: %s", a.sink)
		tctx := cloudevents.ContextWithTarget(context.Background(), a.sink)
		a.ceClient.Send(tctx, csnfCEvent)
		return nil, cloudevents.ResultACK
	} else {
		a.logger.Infof("Sending event to replier")
		return &csnfCEvent, cloudevents.ResultACK
	}
	return &csnfCEvent, cloudevents.ResultACK
}
