// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_homeauto"
	"log"
	"net/http"
	"testing"
)

var x_homeautoMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_homeauto",
	HandleFunc: x_homeautoHandler,
}

func TestX_AVM_DE_Homeauto(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_homeautoMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &x_homeauto.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_Homeauto:1",
			ServiceId:         "urn:X_AVM-DE_Homeauto-com:serviceId:X_AVM-DE_Homeauto1",
			ServiceControlUrl: "/upnp/control/x_homeauto",
		},
	}
	{
		out := &x_homeauto.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_homeauto.GetGenericDeviceInfosRequest{}
		out := &x_homeauto.GetGenericDeviceInfosResponse{}
		require.NoError(t, serviceClient.GetGenericDeviceInfos(in, out))
	}
	{
		in := &x_homeauto.GetSpecificDeviceInfosRequest{}
		out := &x_homeauto.GetSpecificDeviceInfosResponse{}
		require.NoError(t, serviceClient.GetSpecificDeviceInfos(in, out))
	}
	{
		in := &x_homeauto.SetDeviceNameRequest{}
		require.NoError(t, serviceClient.SetDeviceName(in))
	}
	{
		in := &x_homeauto.SetSwitchRequest{}
		require.NoError(t, serviceClient.SetSwitch(in))
	}
}

func x_homeautoHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_homeauto_GetInfo(w)
	case "GetGenericDeviceInfos":
		x_homeauto_GetGenericDeviceInfos(w)
	case "GetSpecificDeviceInfos":
		x_homeauto_GetSpecificDeviceInfos(w)
	case "SetDeviceName":
		x_homeauto_SetDeviceName(w)
	case "SetSwitch":
		x_homeauto_SetSwitch(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_homeauto_GetInfo(w http.ResponseWriter) {
	out := x_homeauto.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_homeauto_GetGenericDeviceInfos(w http.ResponseWriter) {
	out := x_homeauto.GetGenericDeviceInfosResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_homeauto_GetSpecificDeviceInfos(w http.ResponseWriter) {
	out := x_homeauto.GetSpecificDeviceInfosResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_homeauto_SetDeviceName(w http.ResponseWriter) {
	out := x_homeauto.SetDeviceNameResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_homeauto_SetSwitch(w http.ResponseWriter) {
	out := x_homeauto.SetSwitchResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
