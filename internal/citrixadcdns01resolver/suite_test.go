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

package citrixadcdns01resolver

import (
	"flag"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	klog "k8s.io/klog/v2"
)

func init() {
	klog.InitFlags(nil)

	if err := flag.Set("logtostderr", "true"); err != nil {
		os.Exit(-1)
	}

	klog.SetOutput(os.Stdout)
}

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Citrix ADC Certmanager Test Suite")
}
