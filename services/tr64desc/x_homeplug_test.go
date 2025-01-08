// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_homeplug"
	"log"
	"net/http"
	"testing"
)

var x_homeplugMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_homeplug",
	HandleFunc: x_homeplugHandler,
}

func TestX_AVM_DE_Homeplug(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_homeplugMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &x_homeplug.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_Homeplug:1",
			ServiceId:         "urn:X_AVM-DE_Homeplug-com:serviceId:X_AVM-DE_Homeplug1",
			ServiceControlUrl: "/upnp/control/x_homeplug",
		},
	}
	{
		out := &x_homeplug.GetNumberOfDeviceEntriesResponse{}
		require.NoError(t, serviceClient.GetNumberOfDeviceEntries(out))
	}
	{
		in := &x_homeplug.GetGenericDeviceEntryRequest{}
		out := &x_homeplug.GetGenericDeviceEntryResponse{}
		require.NoError(t, serviceClient.GetGenericDeviceEntry(in, out))
	}
	{
		in := &x_homeplug.GetSpecificDeviceEntryRequest{}
		out := &x_homeplug.GetSpecificDeviceEntryResponse{}
		require.NoError(t, serviceClient.GetSpecificDeviceEntry(in, out))
	}
	{
		in := &x_homeplug.DeviceDoUpdateRequest{}
		require.NoError(t, serviceClient.DeviceDoUpdate(in))
	}
}

func x_homeplugHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetNumberOfDeviceEntries":
		x_homeplug_GetNumberOfDeviceEntries(w)
	case "GetGenericDeviceEntry":
		x_homeplug_GetGenericDeviceEntry(w)
	case "GetSpecificDeviceEntry":
		x_homeplug_GetSpecificDeviceEntry(w)
	case "DeviceDoUpdate":
		x_homeplug_DeviceDoUpdate(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_homeplug_GetNumberOfDeviceEntries(w http.ResponseWriter) {
	out := x_homeplug.GetNumberOfDeviceEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_homeplug_GetGenericDeviceEntry(w http.ResponseWriter) {
	out := x_homeplug.GetGenericDeviceEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_homeplug_GetSpecificDeviceEntry(w http.ResponseWriter) {
	out := x_homeplug.GetSpecificDeviceEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_homeplug_DeviceDoUpdate(w http.ResponseWriter) {
	out := x_homeplug.DeviceDoUpdateResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
