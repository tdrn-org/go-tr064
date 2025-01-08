// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_uspcontroller"
	"log"
	"net/http"
	"testing"
)

var x_uspcontrollerMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_uspcontroller",
	HandleFunc: x_uspcontrollerHandler,
}

func TestX_AVM_DE_USPController(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_uspcontrollerMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &x_uspcontroller.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_USPController:1",
			ServiceId:         "urn:X_AVM-DE_USPController-com:serviceId:X_AVM-DE_USPController1",
			ServiceControlUrl: "/upnp/control/x_uspcontroller",
		},
	}
	{
		out := &x_uspcontroller.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_uspcontroller.GetUSPControllerByIndexRequest{}
		out := &x_uspcontroller.GetUSPControllerByIndexResponse{}
		require.NoError(t, serviceClient.GetUSPControllerByIndex(in, out))
	}
	{
		out := &x_uspcontroller.GetUSPControllerNumberOfEntriesResponse{}
		require.NoError(t, serviceClient.GetUSPControllerNumberOfEntries(out))
	}
	{
		in := &x_uspcontroller.AddUSPControllerRequest{}
		out := &x_uspcontroller.AddUSPControllerResponse{}
		require.NoError(t, serviceClient.AddUSPController(in, out))
	}
	{
		in := &x_uspcontroller.DeleteUSPControllerByIndexRequest{}
		require.NoError(t, serviceClient.DeleteUSPControllerByIndex(in))
	}
	{
		in := &x_uspcontroller.SetUSPControllerEnableByIndexRequest{}
		require.NoError(t, serviceClient.SetUSPControllerEnableByIndex(in))
	}
	{
		out := &x_uspcontroller.GetUSPMyFRITZEnableResponse{}
		require.NoError(t, serviceClient.GetUSPMyFRITZEnable(out))
	}
	{
		in := &x_uspcontroller.SetUSPMyFRITZEnableRequest{}
		require.NoError(t, serviceClient.SetUSPMyFRITZEnable(in))
	}
}

func x_uspcontrollerHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_uspcontroller_GetInfo(w)
	case "GetUSPControllerByIndex":
		x_uspcontroller_GetUSPControllerByIndex(w)
	case "GetUSPControllerNumberOfEntries":
		x_uspcontroller_GetUSPControllerNumberOfEntries(w)
	case "AddUSPController":
		x_uspcontroller_AddUSPController(w)
	case "DeleteUSPControllerByIndex":
		x_uspcontroller_DeleteUSPControllerByIndex(w)
	case "SetUSPControllerEnableByIndex":
		x_uspcontroller_SetUSPControllerEnableByIndex(w)
	case "GetUSPMyFRITZEnable":
		x_uspcontroller_GetUSPMyFRITZEnable(w)
	case "SetUSPMyFRITZEnable":
		x_uspcontroller_SetUSPMyFRITZEnable(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_uspcontroller_GetInfo(w http.ResponseWriter) {
	out := x_uspcontroller.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_uspcontroller_GetUSPControllerByIndex(w http.ResponseWriter) {
	out := x_uspcontroller.GetUSPControllerByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_uspcontroller_GetUSPControllerNumberOfEntries(w http.ResponseWriter) {
	out := x_uspcontroller.GetUSPControllerNumberOfEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_uspcontroller_AddUSPController(w http.ResponseWriter) {
	out := x_uspcontroller.AddUSPControllerResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_uspcontroller_DeleteUSPControllerByIndex(w http.ResponseWriter) {
	out := x_uspcontroller.DeleteUSPControllerByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_uspcontroller_SetUSPControllerEnableByIndex(w http.ResponseWriter) {
	out := x_uspcontroller.SetUSPControllerEnableByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_uspcontroller_GetUSPMyFRITZEnable(w http.ResponseWriter) {
	out := x_uspcontroller.GetUSPMyFRITZEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_uspcontroller_SetUSPMyFRITZEnable(w http.ResponseWriter) {
	out := x_uspcontroller.SetUSPMyFRITZEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
