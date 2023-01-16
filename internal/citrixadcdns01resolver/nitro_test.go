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
	"github.com/citrix/adc-nitro-go/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Nitro Operations", func() {
	It("should add txt record", func() {
		client, err := service.NewNitroClientFromEnv()
		Expect(err).To(BeNil(), "cannot loging citrix adc")

		err = nitroAddTxtRecord(client, "_acme-challenge.test.test.kubeservices.io", "random value")
		Expect(err).To(BeNil(), "error at creating txt record")

		res := nitroExistsTxtRecord(client, "_acme-challenge.test.test.kubeservices.io")
		Expect(res).To(BeTrue(), "txt record not exists")

		err = client.Logout()
		Expect(err).To(BeNil(), "error at logout")
	})

	It("should delete txt record", func() {
		client, err := service.NewNitroClientFromEnv()
		Expect(err).To(BeNil(), "cannot loging citrix adc")

		err = nitroAddTxtRecord(client, "_acme-challenge.test.test.kubeservices.io", "random value")
		Expect(err).To(BeNil(), "error at creating txt record")

		res := nitroExistsTxtRecord(client, "_acme-challenge.test.test.kubeservices.io")
		Expect(res).To(BeTrue(), "txt record not exists")

		err = nitroDeleteTxtRecord(client, "_acme-challenge.test.test.kubeservices.io")
		Expect(err).To(BeNil(), "error at deleting txt record")

		res = nitroExistsTxtRecord(client, "_acme-challenge.test.test.kubeservices.io")
		Expect(res).To(BeFalse(), "txt recordt exists")

		err = client.Logout()
		Expect(err).To(BeNil(), "error at logout")
	})
})
