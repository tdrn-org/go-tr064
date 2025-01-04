// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_tam"
	"log"
	"net/http"
	"testing"
)

var x_tamMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_tam",
	HandleFunc: x_tamHandler,
}

func TestX_AVM_DE_TAM(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_tamMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.TR064Spec("tr64desc"))
	client.Debug = true
	serviceClient := &x_tam.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "X_AVM_DE_TAM",
			ServiceType: "urn:dslforum-org:service:X_AVM-DE_TAM:1",
			ServiceId:   "urn:X_AVM-DE_TAM-com:serviceId:X_AVM-DE_TAM1",
			ServiceUrl:  "/upnp/control/x_tam",
		},
	}
	{
		in := &x_tam.GetInfoRequest{}
		out := &x_tam.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(in, out))
	}
	{
		in := &x_tam.SetEnableRequest{}
		require.NoError(t, serviceClient.SetEnable(in))
	}
	{
		in := &x_tam.GetMessageListRequest{}
		out := &x_tam.GetMessageListResponse{}
		require.NoError(t, serviceClient.GetMessageList(in, out))
	}
	{
		in := &x_tam.MarkMessageRequest{}
		require.NoError(t, serviceClient.MarkMessage(in))
	}
	{
		in := &x_tam.DeleteMessageRequest{}
		require.NoError(t, serviceClient.DeleteMessage(in))
	}
	{
		out := &x_tam.GetListResponse{}
		require.NoError(t, serviceClient.GetList(out))
	}
}

func x_tamHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		x_tam_GetInfo(w)
	case "SetEnable":
		x_tam_SetEnable(w)
	case "GetMessageList":
		x_tam_GetMessageList(w)
	case "MarkMessage":
		x_tam_MarkMessage(w)
	case "DeleteMessage":
		x_tam_DeleteMessage(w)
	case "GetList":
		x_tam_GetList(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_tam_GetInfo(w http.ResponseWriter) {
	out := x_tam.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_tam_SetEnable(w http.ResponseWriter) {
	out := x_tam.SetEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_tam_GetMessageList(w http.ResponseWriter) {
	out := x_tam.GetMessageListResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_tam_MarkMessage(w http.ResponseWriter) {
	out := x_tam.MarkMessageResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_tam_DeleteMessage(w http.ResponseWriter) {
	out := x_tam.DeleteMessageResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_tam_GetList(w http.ResponseWriter) {
	out := x_tam.GetListResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
