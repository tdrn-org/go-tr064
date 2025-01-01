// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/x_dect"
	"log"
	"net/http"
	"testing"
)

var x_dectMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_dect",
	HandleFunc: x_dectHandler,
}

func TestX_AVM_DE_Dect(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_dectMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &x_dect.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "X_AVM_DE_Dect",
			ServiceType: "urn:dslforum-org:service:X_AVM-DE_Dect:1",
			ServiceId:   "urn:X_AVM-DE_Dect-com:serviceId:X_AVM-DE_Dect1",
			ServiceUrl:  "/upnp/control/x_dect",
		},
	}
	{
		out := &x_dect.GetNumberOfDectEntriesResponse{}
		require.NoError(t, serviceClient.GetNumberOfDectEntries(out))
	}
	{
		in := &x_dect.GetGenericDectEntryRequest{}
		out := &x_dect.GetGenericDectEntryResponse{}
		require.NoError(t, serviceClient.GetGenericDectEntry(in, out))
	}
	{
		in := &x_dect.GetSpecificDectEntryRequest{}
		out := &x_dect.GetSpecificDectEntryResponse{}
		require.NoError(t, serviceClient.GetSpecificDectEntry(in, out))
	}
	{
		in := &x_dect.DectDoUpdateRequest{}
		require.NoError(t, serviceClient.DectDoUpdate(in))
	}
	{
		out := &x_dect.GetDectListPathResponse{}
		require.NoError(t, serviceClient.GetDectListPath(out))
	}
}

func x_dectHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetNumberOfDectEntries":
		x_dect_GetNumberOfDectEntries(w)
	case "GetGenericDectEntry":
		x_dect_GetGenericDectEntry(w)
	case "GetSpecificDectEntry":
		x_dect_GetSpecificDectEntry(w)
	case "DectDoUpdate":
		x_dect_DectDoUpdate(w)
	case "GetDectListPath":
		x_dect_GetDectListPath(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_dect_GetNumberOfDectEntries(w http.ResponseWriter) {
	out := x_dect.GetNumberOfDectEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_dect_GetGenericDectEntry(w http.ResponseWriter) {
	out := x_dect.GetGenericDectEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_dect_GetSpecificDectEntry(w http.ResponseWriter) {
	out := x_dect.GetSpecificDectEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_dect_DectDoUpdate(w http.ResponseWriter) {
	out := x_dect.DectDoUpdateResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_dect_GetDectListPath(w http.ResponseWriter) {
	out := x_dect.GetDectListPathResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
