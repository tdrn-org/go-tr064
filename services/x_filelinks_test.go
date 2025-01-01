// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/x_filelinks"
	"log"
	"net/http"
	"testing"
)

var x_filelinksMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_filelinks",
	HandleFunc: x_filelinksHandler,
}

func TestX_AVM_DE_Filelinks(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_filelinksMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), mock.User, mock.Password)
	client.Debug = true
	serviceClient := &x_filelinks.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "X_AVM_DE_Filelinks",
			ServiceType: "urn:dslforum-org:service:X_AVM-DE_Filelinks:1",
			ServiceId:   "urn:X_AVM-DE_Filelinks-com:serviceId:X_AVM-DE_Filelinks1",
			ServiceUrl:  "/upnp/control/x_filelinks",
		},
	}
	{
		out := &x_filelinks.GetNumberOfFilelinkEntriesResponse{}
		require.NoError(t, serviceClient.GetNumberOfFilelinkEntries(out))
	}
	{
		in := &x_filelinks.GetGenericFilelinkEntryRequest{}
		out := &x_filelinks.GetGenericFilelinkEntryResponse{}
		require.NoError(t, serviceClient.GetGenericFilelinkEntry(in, out))
	}
	{
		in := &x_filelinks.GetSpecificFilelinkEntryRequest{}
		out := &x_filelinks.GetSpecificFilelinkEntryResponse{}
		require.NoError(t, serviceClient.GetSpecificFilelinkEntry(in, out))
	}
	{
		in := &x_filelinks.NewFilelinkEntryRequest{}
		out := &x_filelinks.NewFilelinkEntryResponse{}
		require.NoError(t, serviceClient.NewFilelinkEntry(in, out))
	}
	{
		in := &x_filelinks.SetFilelinkEntryRequest{}
		require.NoError(t, serviceClient.SetFilelinkEntry(in))
	}
	{
		in := &x_filelinks.DeleteFilelinkEntryRequest{}
		require.NoError(t, serviceClient.DeleteFilelinkEntry(in))
	}
	{
		out := &x_filelinks.GetFilelinkListPathResponse{}
		require.NoError(t, serviceClient.GetFilelinkListPath(out))
	}
}

func x_filelinksHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetNumberOfFilelinkEntries":
		x_filelinks_GetNumberOfFilelinkEntries(w)
	case "GetGenericFilelinkEntry":
		x_filelinks_GetGenericFilelinkEntry(w)
	case "GetSpecificFilelinkEntry":
		x_filelinks_GetSpecificFilelinkEntry(w)
	case "NewFilelinkEntry":
		x_filelinks_NewFilelinkEntry(w)
	case "SetFilelinkEntry":
		x_filelinks_SetFilelinkEntry(w)
	case "DeleteFilelinkEntry":
		x_filelinks_DeleteFilelinkEntry(w)
	case "GetFilelinkListPath":
		x_filelinks_GetFilelinkListPath(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_filelinks_GetNumberOfFilelinkEntries(w http.ResponseWriter) {
	out := x_filelinks.GetNumberOfFilelinkEntriesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_filelinks_GetGenericFilelinkEntry(w http.ResponseWriter) {
	out := x_filelinks.GetGenericFilelinkEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_filelinks_GetSpecificFilelinkEntry(w http.ResponseWriter) {
	out := x_filelinks.GetSpecificFilelinkEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_filelinks_NewFilelinkEntry(w http.ResponseWriter) {
	out := x_filelinks.NewFilelinkEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_filelinks_SetFilelinkEntry(w http.ResponseWriter) {
	out := x_filelinks.SetFilelinkEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_filelinks_DeleteFilelinkEntry(w http.ResponseWriter) {
	out := x_filelinks.DeleteFilelinkEntryResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_filelinks_GetFilelinkListPath(w http.ResponseWriter) {
	out := x_filelinks.GetFilelinkListPathResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
