// generated from spec version: 1.0
package services_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/userif"
)

var userifMock = &mock.ServiceMock{
	Path:       "/upnp/control/userif",
	HandleFunc: userifHandler,
}

func TestUserInterface(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", userifMock)
	defer tr064Mock.Stop(t.Context())
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &userif.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:UserInterface:1",
			ServiceId:         "urn:UserInterface-com:serviceId:UserInterface1",
			ServiceControlUrl: "/upnp/control/userif",
		},
	}
	{
		out := &userif.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		require.NoError(t, serviceClient.X_AVM_DE_CheckUpdate())
	}
	{
		out := &userif.X_AVM_DE_DoUpdateResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_DoUpdate(out))
	}
	{
		out := &userif.X_AVM_DE_DoPrepareCGIResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_DoPrepareCGI(out))
	}
	{
		in := &userif.X_AVM_DE_DoManualUpdateRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_DoManualUpdate(in))
	}
	{
		out := &userif.X_AVM_DE_GetInternationalConfigResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetInternationalConfig(out))
	}
	{
		in := &userif.X_AVM_DE_SetInternationalConfigRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetInternationalConfig(in))
	}
	{
		out := &userif.X_AVM_DE_GetInfoResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetInfo(out))
	}
	{
		in := &userif.X_AVM_DE_SetConfigRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetConfig(in))
	}
}

func userifHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		userif_GetInfo(w)
	case "X_AVM-DE_CheckUpdate":
		userif_X_AVM_DE_CheckUpdate(w)
	case "X_AVM-DE_DoUpdate":
		userif_X_AVM_DE_DoUpdate(w)
	case "X_AVM-DE_DoPrepareCGI":
		userif_X_AVM_DE_DoPrepareCGI(w)
	case "X_AVM-DE_DoManualUpdate":
		userif_X_AVM_DE_DoManualUpdate(w)
	case "X_AVM-DE_GetInternationalConfig":
		userif_X_AVM_DE_GetInternationalConfig(w)
	case "X_AVM-DE_SetInternationalConfig":
		userif_X_AVM_DE_SetInternationalConfig(w)
	case "X_AVM-DE_GetInfo":
		userif_X_AVM_DE_GetInfo(w)
	case "X_AVM-DE_SetConfig":
		userif_X_AVM_DE_SetConfig(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func userif_GetInfo(w http.ResponseWriter) {
	out := userif.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func userif_X_AVM_DE_CheckUpdate(w http.ResponseWriter) {
	out := userif.X_AVM_DE_CheckUpdateResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func userif_X_AVM_DE_DoUpdate(w http.ResponseWriter) {
	out := userif.X_AVM_DE_DoUpdateResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func userif_X_AVM_DE_DoPrepareCGI(w http.ResponseWriter) {
	out := userif.X_AVM_DE_DoPrepareCGIResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func userif_X_AVM_DE_DoManualUpdate(w http.ResponseWriter) {
	out := userif.X_AVM_DE_DoManualUpdateResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func userif_X_AVM_DE_GetInternationalConfig(w http.ResponseWriter) {
	out := userif.X_AVM_DE_GetInternationalConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func userif_X_AVM_DE_SetInternationalConfig(w http.ResponseWriter) {
	out := userif.X_AVM_DE_SetInternationalConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func userif_X_AVM_DE_GetInfo(w http.ResponseWriter) {
	out := userif.X_AVM_DE_GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func userif_X_AVM_DE_SetConfig(w http.ResponseWriter) {
	out := userif.X_AVM_DE_SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
