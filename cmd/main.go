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
	"fmt"
	"os"

	webhookcmd "github.com/cert-manager/cert-manager/pkg/acme/webhook/cmd"
	dr "github.com/kazimsarikaya/citrixadccertmanagerwebhook/internal/citrixadcdns01resolver"
	klog "k8s.io/klog/v2"
)

var (
	GroupName = os.Getenv("GROUP_NAME")
	version   = ""
	buildTime = ""
	goVersion = ""
)

func init() {
	klog.InitFlags(nil)
	klog.SetOutput(os.Stdout)
}

func printVersion() {
	fmt.Printf("Citrix DNS01 cert manager webhook\n")
	fmt.Printf("Version: %v\n", version)
	fmt.Printf("Build Time: %v\n", buildTime)
	fmt.Printf("%v\n", goVersion)
}

func main() {
	if GroupName == "" {
		klog.Errorf("GROUP_NAME must be specified")

		os.Exit(-1)
	}

	printVersion()

	// This will register our custom DNS provider with the webhook serving
	// library, making it available as an API under the provided GroupName.
	// You can register multiple DNS provider implementations with a single
	// webhook, where the Name() method will be used to disambiguate between
	// the different implementations.
	webhookcmd.RunWebhookServer(GroupName,
		dr.NewCitrixAdcDNS01Resolver(),
	)
}
