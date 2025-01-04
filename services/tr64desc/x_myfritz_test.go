// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_myfritz"
	"log"
	"net/http"
	"testing"
)

var x_myfritzMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_myfritz",
	HandleFunc: x_myfritzHandler,
}

func TestX_AVM_DE_MyFritz(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_myfritzMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &x_myfritz.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec: tr064.ServiceSpec("tr64desc"),
			ServiceId:   "urn:X_AVM-DE_MyFritz-com:serviceId:X_AVM-DE_MyFritz1",
			ServiceUrl:  "/upnp/control/x_myfritz",
		},
	}
	{
		out := &x_myfritz.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_myfritz.SetMyFRITZRequest{}
		require.NoError(t, serviceClient.SetMyFRITZ(in))
	}
	{
		out := &x_myfritz.GetNumberOfServicesResponse{}
		require.NoError(t, serviceClient.GetNumberOfServices(out))
	}
	{
		in := &x_myfritz.GetServiceByIndexRequest{}
		out := &x_myfritz.GetServiceByIndexResponse{}
		require.NoError(t, serviceClient.GetServiceByIndex(in, out))
	}
	{
		in := &x_myfritz.SetServiceByIndexRequest{}
		require.NoError(t, serviceClient.SetServiceByIndex(in))
	}
	{
		in := &x_myfritz.DeleteServiceByIndexRequest{}
		require.NoError(t, serviceClient.DeleteServiceByIndex(in))
	}
}

func x_myfritzHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_myfritz_GetInfo(w)
	case "SetMyFRITZ":
		x_myfritz_SetMyFRITZ(w)
	case "GetNumberOfServices":
		x_myfritz_GetNumberOfServices(w)
	case "GetServiceByIndex":
		x_myfritz_GetServiceByIndex(w)
	case "SetServiceByIndex":
		x_myfritz_SetServiceByIndex(w)
	case "DeleteServiceByIndex":
		x_myfritz_DeleteServiceByIndex(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_myfritz_GetInfo(w http.ResponseWriter) {
	out := x_myfritz.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_myfritz_SetMyFRITZ(w http.ResponseWriter) {
	out := x_myfritz.SetMyFRITZResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_myfritz_GetNumberOfServices(w http.ResponseWriter) {
	out := x_myfritz.GetNumberOfServicesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_myfritz_GetServiceByIndex(w http.ResponseWriter) {
	out := x_myfritz.GetServiceByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_myfritz_SetServiceByIndex(w http.ResponseWriter) {
	out := x_myfritz.SetServiceByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_myfritz_DeleteServiceByIndex(w http.ResponseWriter) {
	out := x_myfritz.DeleteServiceByIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
