/*
Copyright 2021 TriggerMesh Inc.

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

package main

import (
	aquasectransformation "github.com/triggermesh/CSNFQuickstart/src/aquasectransformation/pkg/adapter"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

func main() {
	pkgadapter.Main("aquasectransformation-adapter", aquasectransformation.EnvAccessorCtor, aquasectransformation.NewAdapter)
}
