// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/wandsllinkconfig"
	"log"
	"net/http"
	"testing"
)

var wandsllinkconfigMock = &mock.ServiceMock{
	Path:       "/upnp/control/wandsllinkconfig1",
	HandleFunc: wandsllinkconfigHandler,
}

func TestWANDSLLinkConfig(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", wandsllinkconfigMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &wandsllinkconfig.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "WANDSLLinkConfig",
			ServiceType: "urn:dslforum-org:service:WANDSLLinkConfig:1",
			ServiceId:   "urn:WANDSLLinkConfig-com:serviceId:WANDSLLinkConfig1",
			ServiceUrl:  "/upnp/control/wandsllinkconfig1",
		},
	}
	{
		out := &wandsllinkconfig.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &wandsllinkconfig.SetEnableRequest{}
		require.NoError(t, serviceClient.SetEnable(in))
	}
	{
		out := &wandsllinkconfig.GetAutoConfigResponse{}
		require.NoError(t, serviceClient.GetAutoConfig(out))
	}
	{
		in := &wandsllinkconfig.SetDSLLinkTypeRequest{}
		require.NoError(t, serviceClient.SetDSLLinkType(in))
	}
	{
		out := &wandsllinkconfig.GetDSLLinkInfoResponse{}
		require.NoError(t, serviceClient.GetDSLLinkInfo(out))
	}
	{
		in := &wandsllinkconfig.SetDestinationAddressRequest{}
		require.NoError(t, serviceClient.SetDestinationAddress(in))
	}
	{
		out := &wandsllinkconfig.GetDestinationAddressResponse{}
		require.NoError(t, serviceClient.GetDestinationAddress(out))
	}
	{
		in := &wandsllinkconfig.SetATMEncapsulationRequest{}
		require.NoError(t, serviceClient.SetATMEncapsulation(in))
	}
	{
		out := &wandsllinkconfig.GetATMEncapsulationResponse{}
		require.NoError(t, serviceClient.GetATMEncapsulation(out))
	}
	{
		out := &wandsllinkconfig.GetStatisticsResponse{}
		require.NoError(t, serviceClient.GetStatistics(out))
	}
}

func wandsllinkconfigHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		wandsllinkconfig_GetInfo(w)
	case "SetEnable":
		wandsllinkconfig_SetEnable(w)
	case "GetAutoConfig":
		wandsllinkconfig_GetAutoConfig(w)
	case "SetDSLLinkType":
		wandsllinkconfig_SetDSLLinkType(w)
	case "GetDSLLinkInfo":
		wandsllinkconfig_GetDSLLinkInfo(w)
	case "SetDestinationAddress":
		wandsllinkconfig_SetDestinationAddress(w)
	case "GetDestinationAddress":
		wandsllinkconfig_GetDestinationAddress(w)
	case "SetATMEncapsulation":
		wandsllinkconfig_SetATMEncapsulation(w)
	case "GetATMEncapsulation":
		wandsllinkconfig_GetATMEncapsulation(w)
	case "GetStatistics":
		wandsllinkconfig_GetStatistics(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func wandsllinkconfig_GetInfo(w http.ResponseWriter) {
	out := wandsllinkconfig.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_SetEnable(w http.ResponseWriter) {
	out := wandsllinkconfig.SetEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_GetAutoConfig(w http.ResponseWriter) {
	out := wandsllinkconfig.GetAutoConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_SetDSLLinkType(w http.ResponseWriter) {
	out := wandsllinkconfig.SetDSLLinkTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_GetDSLLinkInfo(w http.ResponseWriter) {
	out := wandsllinkconfig.GetDSLLinkInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_SetDestinationAddress(w http.ResponseWriter) {
	out := wandsllinkconfig.SetDestinationAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_GetDestinationAddress(w http.ResponseWriter) {
	out := wandsllinkconfig.GetDestinationAddressResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_SetATMEncapsulation(w http.ResponseWriter) {
	out := wandsllinkconfig.SetATMEncapsulationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_GetATMEncapsulation(w http.ResponseWriter) {
	out := wandsllinkconfig.GetATMEncapsulationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wandsllinkconfig_GetStatistics(w http.ResponseWriter) {
	out := wandsllinkconfig.GetStatisticsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
