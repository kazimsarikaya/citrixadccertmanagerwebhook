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
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/citrix/adc-nitro-go/service"
	"k8s.io/client-go/rest"
	klog "k8s.io/klog/v2"
)

const webhookName = "citrix-adc-cert-manager-webhook"

type citrixAdcDNS01Resolver struct {
	client *service.NitroClient
}

func NewCitrixAdcDNS01Resolver() *citrixAdcDNS01Resolver {
	return &citrixAdcDNS01Resolver{}
}

func (c *citrixAdcDNS01Resolver) Name() string {
	return webhookName
}

func (c *citrixAdcDNS01Resolver) Present(ch *v1alpha1.ChallengeRequest) error {
	return nitroAddTxtRecord(c.client, ch.ResolvedFQDN, ch.Key)
}

func (c *citrixAdcDNS01Resolver) CleanUp(ch *v1alpha1.ChallengeRequest) error {
	return nitroDeleteTxtRecord(c.client, ch.ResolvedFQDN)
}

func (c *citrixAdcDNS01Resolver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	client, err := service.NewNitroClientFromEnv()
	if err != nil {
		klog.Errorf("cannot init nitro client: %v", err)

		return err
	}

	c.client = client

	return nil
}
