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
	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"
	klog "k8s.io/klog/v2"
)

func nitroAddTxtRecord(client *service.NitroClient, domain, value string) error {
	if nitroExistsTxtRecord(client, domain) {
		return nil
	}

	txtRec := &dns.Dnstxtrec{
		Domain: domain,
		String: []string{value},
	}

	_, err := client.AddResource(service.Dnstxtrec.Type(), domain, &txtRec)
	if err != nil {
		klog.Errorf("cannot add txt record: %v", err)

		return err
	}

	return nil
}

func nitroExistsTxtRecord(client *service.NitroClient, domain string) bool {
	return client.ResourceExists(service.Dnstxtrec.Type(), domain)
}

func nitroDeleteTxtRecord(client *service.NitroClient, domain string) error {
	if !nitroExistsTxtRecord(client, domain) {
		return nil
	}

	txtRec, err := client.FindResource(service.Dnstxtrec.Type(), domain)
	if err != nil {
		klog.Errorf("cannot find resource %s", err)

		return err
	}

	args := make(map[string]string)

	args["recordid"] = txtRec["recordid"].(string)

	err = client.DeleteResourceWithArgsMap(service.Dnstxtrec.Type(), domain, args)

	if err != nil {
		klog.Errorf("cannot delete txt record: %v", err)

		return err
	}

	return nil
}
