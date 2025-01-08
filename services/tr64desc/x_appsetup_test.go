// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_appsetup"
	"log"
	"net/http"
	"testing"
)

var x_appsetupMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_appsetup",
	HandleFunc: x_appsetupHandler,
}

func TestX_AVM_DE_AppSetup(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_appsetupMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &x_appsetup.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_AVM-DE_AppSetup:1",
			ServiceId:         "urn:X_AVM-DE_AppSetup-com:serviceId:X_AVM-DE_AppSetup1",
			ServiceControlUrl: "/upnp/control/x_appsetup",
		},
	}
	{
		out := &x_appsetup.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		out := &x_appsetup.GetConfigResponse{}
		require.NoError(t, serviceClient.GetConfig(out))
	}
	{
		in := &x_appsetup.GetAppMessageFilterRequest{}
		out := &x_appsetup.GetAppMessageFilterResponse{}
		require.NoError(t, serviceClient.GetAppMessageFilter(in, out))
	}
	{
		in := &x_appsetup.RegisterAppRequest{}
		require.NoError(t, serviceClient.RegisterApp(in))
	}
	{
		in := &x_appsetup.SetAppVPNRequest{}
		require.NoError(t, serviceClient.SetAppVPN(in))
	}
	{
		in := &x_appsetup.SetAppVPNwithPFSRequest{}
		require.NoError(t, serviceClient.SetAppVPNwithPFS(in))
	}
	{
		in := &x_appsetup.SetAppMessageFilterRequest{}
		require.NoError(t, serviceClient.SetAppMessageFilter(in))
	}
	{
		in := &x_appsetup.SetAppMessageReceiverRequest{}
		out := &x_appsetup.SetAppMessageReceiverResponse{}
		require.NoError(t, serviceClient.SetAppMessageReceiver(in, out))
	}
	{
		in := &x_appsetup.ResetEventRequest{}
		require.NoError(t, serviceClient.ResetEvent(in))
	}
	{
		out := &x_appsetup.GetAppRemoteInfoResponse{}
		require.NoError(t, serviceClient.GetAppRemoteInfo(out))
	}
	{
		in := &x_appsetup.GetBoxSenderIdRequest{}
		out := &x_appsetup.GetBoxSenderIdResponse{}
		require.NoError(t, serviceClient.GetBoxSenderId(in, out))
	}
}

func x_appsetupHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_appsetup_GetInfo(w)
	case "GetConfig":
		x_appsetup_GetConfig(w)
	case "GetAppMessageFilter":
		x_appsetup_GetAppMessageFilter(w)
	case "RegisterApp":
		x_appsetup_RegisterApp(w)
	case "SetAppVPN":
		x_appsetup_SetAppVPN(w)
	case "SetAppVPNwithPFS":
		x_appsetup_SetAppVPNwithPFS(w)
	case "SetAppMessageFilter":
		x_appsetup_SetAppMessageFilter(w)
	case "SetAppMessageReceiver":
		x_appsetup_SetAppMessageReceiver(w)
	case "ResetEvent":
		x_appsetup_ResetEvent(w)
	case "GetAppRemoteInfo":
		x_appsetup_GetAppRemoteInfo(w)
	case "GetBoxSenderId":
		x_appsetup_GetBoxSenderId(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_appsetup_GetInfo(w http.ResponseWriter) {
	out := x_appsetup.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_GetConfig(w http.ResponseWriter) {
	out := x_appsetup.GetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_GetAppMessageFilter(w http.ResponseWriter) {
	out := x_appsetup.GetAppMessageFilterResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_RegisterApp(w http.ResponseWriter) {
	out := x_appsetup.RegisterAppResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_SetAppVPN(w http.ResponseWriter) {
	out := x_appsetup.SetAppVPNResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_SetAppVPNwithPFS(w http.ResponseWriter) {
	out := x_appsetup.SetAppVPNwithPFSResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_SetAppMessageFilter(w http.ResponseWriter) {
	out := x_appsetup.SetAppMessageFilterResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_SetAppMessageReceiver(w http.ResponseWriter) {
	out := x_appsetup.SetAppMessageReceiverResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_ResetEvent(w http.ResponseWriter) {
	out := x_appsetup.ResetEventResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_GetAppRemoteInfo(w http.ResponseWriter) {
	out := x_appsetup.GetAppRemoteInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_appsetup_GetBoxSenderId(w http.ResponseWriter) {
	out := x_appsetup.GetBoxSenderIdResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
