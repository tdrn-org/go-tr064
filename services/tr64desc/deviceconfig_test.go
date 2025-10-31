// generated from spec version: 1.0
package services_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/deviceconfig"
)

var deviceconfigMock = &mock.ServiceMock{
	Path:       "/upnp/control/deviceconfig",
	HandleFunc: deviceconfigHandler,
}

func TestDeviceConfig(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", deviceconfigMock)
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &deviceconfig.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:DeviceConfig:1",
			ServiceId:         "urn:DeviceConfig-com:serviceId:DeviceConfig1",
			ServiceControlUrl: "/upnp/control/deviceconfig",
		},
	}
	{
		out := &deviceconfig.GetPersistentDataResponse{}
		require.NoError(t, serviceClient.GetPersistentData(out))
	}
	{
		in := &deviceconfig.SetPersistentDataRequest{}
		require.NoError(t, serviceClient.SetPersistentData(in))
	}
	{
		in := &deviceconfig.ConfigurationStartedRequest{}
		require.NoError(t, serviceClient.ConfigurationStarted(in))
	}
	{
		out := &deviceconfig.ConfigurationFinishedResponse{}
		require.NoError(t, serviceClient.ConfigurationFinished(out))
	}
	{
		require.NoError(t, serviceClient.FactoryReset())
	}
	{
		require.NoError(t, serviceClient.Reboot())
	}
	{
		out := &deviceconfig.X_GenerateUUIDResponse{}
		require.NoError(t, serviceClient.X_GenerateUUID(out))
	}
	{
		in := &deviceconfig.X_AVM_DE_GetConfigFileRequest{}
		out := &deviceconfig.X_AVM_DE_GetConfigFileResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetConfigFile(in, out))
	}
	{
		in := &deviceconfig.X_AVM_DE_SetConfigFileRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetConfigFile(in))
	}
	{
		out := &deviceconfig.X_AVM_DE_CreateUrlSIDResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_CreateUrlSID(out))
	}
	{
		in := &deviceconfig.X_AVM_DE_SendSupportDataRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SendSupportData(in))
	}
	{
		out := &deviceconfig.X_AVM_DE_GetSupportDataInfoResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetSupportDataInfo(out))
	}
	{
		out := &deviceconfig.X_AVM_DE_GetSupportDataEnableResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetSupportDataEnable(out))
	}
	{
		in := &deviceconfig.X_AVM_DE_SetSupportDataEnableRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetSupportDataEnable(in))
	}
}

func deviceconfigHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetPersistentData":
		deviceconfig_GetPersistentData(w)
	case "SetPersistentData":
		deviceconfig_SetPersistentData(w)
	case "ConfigurationStarted":
		deviceconfig_ConfigurationStarted(w)
	case "ConfigurationFinished":
		deviceconfig_ConfigurationFinished(w)
	case "FactoryReset":
		deviceconfig_FactoryReset(w)
	case "Reboot":
		deviceconfig_Reboot(w)
	case "X_GenerateUUID":
		deviceconfig_X_GenerateUUID(w)
	case "X_AVM-DE_GetConfigFile":
		deviceconfig_X_AVM_DE_GetConfigFile(w)
	case "X_AVM-DE_SetConfigFile":
		deviceconfig_X_AVM_DE_SetConfigFile(w)
	case "X_AVM-DE_CreateUrlSID":
		deviceconfig_X_AVM_DE_CreateUrlSID(w)
	case "X_AVM-DE_SendSupportData":
		deviceconfig_X_AVM_DE_SendSupportData(w)
	case "X_AVM-DE_GetSupportDataInfo":
		deviceconfig_X_AVM_DE_GetSupportDataInfo(w)
	case "X_AVM-DE_GetSupportDataEnable":
		deviceconfig_X_AVM_DE_GetSupportDataEnable(w)
	case "X_AVM-DE_SetSupportDataEnable":
		deviceconfig_X_AVM_DE_SetSupportDataEnable(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deviceconfig_GetPersistentData(w http.ResponseWriter) {
	out := deviceconfig.GetPersistentDataResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_SetPersistentData(w http.ResponseWriter) {
	out := deviceconfig.SetPersistentDataResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_ConfigurationStarted(w http.ResponseWriter) {
	out := deviceconfig.ConfigurationStartedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_ConfigurationFinished(w http.ResponseWriter) {
	out := deviceconfig.ConfigurationFinishedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_FactoryReset(w http.ResponseWriter) {
	out := deviceconfig.FactoryResetResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_Reboot(w http.ResponseWriter) {
	out := deviceconfig.RebootResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_X_GenerateUUID(w http.ResponseWriter) {
	out := deviceconfig.X_GenerateUUIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_X_AVM_DE_GetConfigFile(w http.ResponseWriter) {
	out := deviceconfig.X_AVM_DE_GetConfigFileResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_X_AVM_DE_SetConfigFile(w http.ResponseWriter) {
	out := deviceconfig.X_AVM_DE_SetConfigFileResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_X_AVM_DE_CreateUrlSID(w http.ResponseWriter) {
	out := deviceconfig.X_AVM_DE_CreateUrlSIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_X_AVM_DE_SendSupportData(w http.ResponseWriter) {
	out := deviceconfig.X_AVM_DE_SendSupportDataResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_X_AVM_DE_GetSupportDataInfo(w http.ResponseWriter) {
	out := deviceconfig.X_AVM_DE_GetSupportDataInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_X_AVM_DE_GetSupportDataEnable(w http.ResponseWriter) {
	out := deviceconfig.X_AVM_DE_GetSupportDataEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceconfig_X_AVM_DE_SetSupportDataEnable(w http.ResponseWriter) {
	out := deviceconfig.X_AVM_DE_SetSupportDataEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
