// Copyright 2024 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package features

import (
	"github.com/crossplane/crossplane-runtime/pkg/feature"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
)

const (
	// EnableAlphaSharedTelemetry enables alpha support for Telemetry.
	EnableAlphaSharedTelemetry feature.Flag = "EnableSharedTelemetry"
)

func EnableFeatures(features *feature.Flags, params map[string]any) {
	if isAlphaSharedTelemetryEnabled(params) {
		features.Enable(EnableAlphaSharedTelemetry)
	}
}

func isAlphaSharedTelemetryEnabled(params map[string]any) bool {
	enabled, err := fieldpath.Pave(params).GetBool("features.alpha.observability.enabled")
	if err != nil {
		return false
	}
	return enabled
}
