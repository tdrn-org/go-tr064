// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/layer3forwarding"
	"log"
	"net/http"
	"testing"
)

var layer3forwardingMock = &mock.ServiceMock{
	Path:       "/upnp/control/layer3forwarding",
	HandleFunc: layer3forwardingHandler,
}

func TestLayer3Forwarding(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", layer3forwardingMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &layer3forwarding.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:Layer3Forwarding:1",
			ServiceId:         "urn:Layer3Forwarding-com:serviceId:Layer3Forwarding1",
			ServiceControlUrl: "/upnp/control/layer3forwarding",
		},
	}
	{
		in := &layer3forwarding.SetDefaultConnectionServiceRequest{}
		require.NoError(t, serviceClient.SetDefaultConnectionService(in))
	}
	{
		out := &layer3forwarding.GetDefaultConnectionServiceResponse{}
		require.NoError(t, serviceClient.GetDefaultConnectionService(out))
	}
	{
		out := &layer3forwarding.GetForwardNumberOfEntriesResponse{}
		require.NoError(t, serviceClient.GetForwardNumberOfEntries(out))
	}
	{
		in := &layer3forwarding.AddForwardingEntryRequest{}
		require.NoError(t, serviceClient.AddForwardingEntry(in))
	}
	{
		in := &layer3forwarding.DeleteForwardingEntryRequest{}
		require.NoError(t, serviceClient.DeleteForwardingEntry(in))
	}
	{
		in := &layer3forwarding.GetSpecificForwardingEntryRequest{}
		out := &layer3forwarding.GetSpecificForwardingEntryResponse{}
		require.NoError(t, serviceClient.GetSpecificForwardingEntry(in, out))
	}
	{
		in := &layer3forwarding.GetGenericForwardingEntryRequest{}
		out := &layer3forwarding.GetGenericForwardingEntryResponse{}
		require.NoError(t, serviceClient.GetGenericForwardingEntry(in, out))
	}
	{
		in := &layer3forwarding.SetForwardingEntryEnableRequest{}
		require.NoError(t, serviceClient.SetForwardingEntryEnable(in))
	}
}

func layer3forwardingHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "SetDefaultConnectionService":
		layer3forwarding_SetDefaultConnectionService(w)
	case "GetDefaultConnectionService":
		layer3forwarding_GetDefaultConnectionService(w)
	case "GetForwardNumberOfEntries":
		layer3forwarding_GetForwardNumberOfEntries(w)
	case "AddForwardingEntry":
		layer3forwarding_AddForwardingEntry(w)
	case "DeleteForwardingEntry":
		layer3forwarding_DeleteForwardingEntry(w)
	case "GetSpecificForwardingEntry":
		layer3forwarding_GetSpecificForwardingEntry(w)
	case "GetGenericForwardingEntry":
		layer3forwarding_GetGenericForwardingEntry(w)
	case "SetForwardingEntryEnable":
		layer3forwarding_SetForwardingEntryEnable(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func layer3forwarding_SetDefaultConnectionService(w http.ResponseWriter) {
	out := layer3forwarding.SetDefaultConnectionServiceResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func layer3forwarding_GetDefaultConnectionService(w http.ResponseWriter) {
	out := layer3forwarding.GetDefaultConnectionServiceResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func layer3forwarding_GetForwardNumberOfEntries(w http.ResponseWriter) {
	out := layer3forwarding.GetForwardNumberOfEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func layer3forwarding_AddForwardingEntry(w http.ResponseWriter) {
	out := layer3forwarding.AddForwardingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func layer3forwarding_DeleteForwardingEntry(w http.ResponseWriter) {
	out := layer3forwarding.DeleteForwardingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func layer3forwarding_GetSpecificForwardingEntry(w http.ResponseWriter) {
	out := layer3forwarding.GetSpecificForwardingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func layer3forwarding_GetGenericForwardingEntry(w http.ResponseWriter) {
	out := layer3forwarding.GetGenericForwardingEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func layer3forwarding_SetForwardingEntryEnable(w http.ResponseWriter) {
	out := layer3forwarding.SetForwardingEntryEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
