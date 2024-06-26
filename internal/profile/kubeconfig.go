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

package profile

import (
	"context"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

// GetIngressHost returns the ingress host of the Spaces cfg points to. If the
// ingress is not configured, it returns an empty string.
func GetIngressHost(ctx context.Context, cl corev1client.ConfigMapsGetter) (host string, ca []byte, err error) {
	mxpConfig, err := cl.ConfigMaps("upbound-system").Get(ctx, "ingress-public", metav1.GetOptions{})
	if err != nil {
		return "", nil, err
	}

	host = mxpConfig.Data["ingress-host"]
	ca = []byte(mxpConfig.Data["ingress-ca"])
	return strings.TrimPrefix(host, "https://"), ca, nil
}
