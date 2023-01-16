/*
Copyright 2023 Kazım SARIKAYA

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
	"os"

	dr "github.com/kazimsarikaya/citrixadccertmanagerwebhook/internal/citrixadcdns01resolver" 
	webhookcmd "github.com/cert-manager/cert-manager/pkg/acme/webhook/cmd"
	klog "k8s.io/klog/v2"
)

var GroupName = os.Getenv("GROUP_NAME")

func main() {
	if GroupName == "" {
		klog.Errorf("GROUP_NAME must be specified")

		os.Exit(-1)
	}

	// This will register our custom DNS provider with the webhook serving
	// library, making it available as an API under the provided GroupName.
	// You can register multiple DNS provider implementations with a single
	// webhook, where the Name() method will be used to disambiguate between
	// the different implementations.
	webhookcmd.RunWebhookServer(GroupName,
		dr.NewCitrixAdcDNS01Resolver(),
	)
}
