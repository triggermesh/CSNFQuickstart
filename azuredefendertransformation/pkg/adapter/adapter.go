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

// Package AZUREDEFENDERTRANSFORMATION implements a CloudEvents adapter that...
package azuredefendertransformation

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
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		VendorName         string    `json:"vendorName"`
		AlertDisplayName   string    `json:"alertDisplayName"`
		AlertName          string    `json:"alertName"`
		DetectedTimeUtc    time.Time `json:"detectedTimeUtc"`
		Description        string    `json:"description"`
		RemediationSteps   string    `json:"remediationSteps"`
		ActionTaken        string    `json:"actionTaken"`
		ReportedSeverity   string    `json:"reportedSeverity"`
		CompromisedEntity  string    `json:"compromisedEntity"`
		AssociatedResource string    `json:"associatedResource"`
		SubscriptionID     string    `json:"subscriptionId"`
		InstanceID         string    `json:"instanceId"`
		ExtendedProperties struct {
			ResourceType               string `json:"resourceType"`
			PotentialCauses            string `json:"potential causes"`
			ClientPrincipalName        string `json:"client principal name"`
			AlertID                    string `json:"alert Id"`
			ClientIPAddress            string `json:"client IP address"`
			ClientIPLocation           string `json:"client IP location"`
			ClientApplication          string `json:"client application"`
			SuccessfulLogins           string `json:"successful logins"`
			OmsWorkspaceID             string `json:"oms workspace ID"`
			FailedLogins               string `json:"failed logins"`
			OmsAgentID                 string `json:"oms agent ID"`
			EnrichmentTasThreatReports string `json:"enrichment_tas_threat__reports"`
			KillChainIntent            string `json:"killChainIntent"`
		} `json:"extendedProperties"`
		State             string        `json:"state"`
		ReportedTimeUtc   time.Time     `json:"reportedTimeUtc"`
		ConfidenceReasons []interface{} `json:"confidenceReasons"`
		CanBeInvestigated bool          `json:"canBeInvestigated"`
		IsIncident        bool          `json:"isIncident"`
		Entities          []struct {
			ID         string `json:"$id"`
			HostName   string `json:"hostName,omitempty"`
			AzureID    string `json:"azureID,omitempty"`
			OmsAgentID string `json:"omsAgentID,omitempty"`
			Type       string `json:"type"`
			Address    string `json:"address,omitempty"`
			Location   struct {
				CountryCode      string `json:"countryCode"`
				CountryName      string `json:"countryName"`
				State            string `json:"state"`
				City             string `json:"city"`
				Longitude        int    `json:"longitude"`
				Latitude         int    `json:"latitude"`
				Asn              int    `json:"asn"`
				Carrier          string `json:"carrier"`
				Organization     string `json:"organization"`
				OrganizationType string `json:"organizationType"`
				CloudProvider    string `json:"cloudProvider"`
				SystemService    string `json:"systemService"`
			} `json:"location,omitempty"`
			SourceAddress struct {
				Ref string `json:"$ref"`
			} `json:"sourceAddress,omitempty"`
			Protocol string `json:"protocol,omitempty"`
			Name     string `json:"name,omitempty"`
			Host     struct {
				Ref string `json:"$ref"`
			} `json:"host,omitempty"`
		} `json:"entities"`
	} `json:"properties"`
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
		targetce.ReplierWithStaticResponseType("io.triggermesh.azuredefendertransformation.error"),
		targetce.ReplierWithPayloadPolicy(targetce.PayloadPolicy(env.CloudEventPayloadPolicy)))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	return &azuredefendertransformationadapter{
		sink:     env.Sink,
		replier:  replier,
		ceClient: ceClient,
		logger:   logger,
	}
}

var _ pkgadapter.Adapter = (*azuredefendertransformationadapter)(nil)

type azuredefendertransformationadapter struct {
	sink     string
	replier  *targetce.Replier
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger
}

// Start is a blocking function and will return if an error occurs
// or the context is cancelled.
func (a *azuredefendertransformationadapter) Start(ctx context.Context) error {
	a.logger.Info("Starting AZUREDEFENDERTRANSFORMATION Adapter")
	return a.ceClient.StartReceiver(ctx, a.dispatch)
}

func (a *azuredefendertransformationadapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
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
			GUID:             ade.ID,
			Name:             ade.Name,
			Severity:         ade.Properties.ExtendedProperties.KillChainIntent,
			ShortDescription: ade.Properties.Description,
			StartTime:        ade.Properties.ReportedTimeUtc,
			Status:           ade.Properties.State,
		},
		Provider: struct {
			AccountID string `json:"accountId"`
		}{
			AccountID: ade.ID,
		},
		ProviderID:   ade.ID,
		ProviderType: "Azure",
		Resource: struct {
			Identifier string `json:"identifier"`
			Name       string `json:"name"`
			Region     string `json:"region"`
			Type       string `json:"type"`
			Zone       string `json:"zone"`
		}{
			Identifier: ade.Properties.InstanceID,
			Name:       ade.Properties.AlertName,
			Region:     "us-east-1",
			Type:       ade.Type,
			Zone:       "US",
		},
		Source: struct {
			SourceID   string `json:"sourceId"`
			SourceName string `json:"sourceName"`
		}{
			SourceID:   "Azure Defender",
			SourceName: "Azure Defender",
		},
	}

	csnfCEvent := cloudevents.NewEvent()
	csnfCEvent.SetID(csnfEvent.Event.GUID)
	csnfCEvent.SetType("io.triggermesh.csnf.azure.defender.event")
	csnfCEvent.SetSource("AzureDefender")
	csnfCEvent.SetSubject("Azure Defender")
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
