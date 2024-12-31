/*
 * Copyright 2024 Holger de Carne
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mock

import (
	"log"
	"net/http"

	"github.com/tdrn-org/go-tr064/services/deviceconfig"
)

func (mock *mockServer) handleDeviceconfig(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.unmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "ConfigurationFinished":
		mock.writeDeviceconfigConfigurationFinished(w)
	case "ConfigurationStarted":
		mock.writeDeviceconfigConfigurationStarted(w)
	case "FactoryReset":
		mock.writeDeviceconfigFactoryReset(w)
	case "GetPersistentData":
		mock.writeDeviceconfigGetPersistentData(w)
	case "Reboot":
		mock.writeDeviceconfigReboot(w)
	case "SetPersistentData":
		mock.writeDeviceconfigSetPersistentData(w)
	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (mock *mockServer) writeDeviceconfigConfigurationFinished(w http.ResponseWriter) {
	out := &deviceconfig.ConfigurationFinishedResponse{
		NewStatus: "NewStatus",
	}
	err := mock.sendSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func (mock *mockServer) writeDeviceconfigConfigurationStarted(w http.ResponseWriter) {
	out := &deviceconfig.ConfigurationStartedResponse{}
	err := mock.sendSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func (mock *mockServer) writeDeviceconfigFactoryReset(w http.ResponseWriter) {
	out := &deviceconfig.FactoryResetResponse{}
	err := mock.sendSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func (mock *mockServer) writeDeviceconfigGetPersistentData(w http.ResponseWriter) {
	out := &deviceconfig.GetPersistentDataResponse{
		NewPersistentData: "NewPersistentData",
	}
	err := mock.sendSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func (mock *mockServer) writeDeviceconfigReboot(w http.ResponseWriter) {
	out := &deviceconfig.RebootResponse{}
	err := mock.sendSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func (mock *mockServer) writeDeviceconfigSetPersistentData(w http.ResponseWriter) {
	out := &deviceconfig.SetPersistentDataResponse{}
	err := mock.sendSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
