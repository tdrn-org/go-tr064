// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/time"
	"log"
	"net/http"
	"testing"
)

var timeMock = &mock.ServiceMock{
	Path:       "/upnp/control/time",
	HandleFunc: timeHandler,
}

func TestTime(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", timeMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &time.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "Time",
			ServiceType: "urn:dslforum-org:service:Time:1",
			ServiceId:   "urn:Time-com:serviceId:Time1",
			ServiceUrl:  "/upnp/control/time",
		},
	}
	{
		out := &time.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &time.SetNTPServersRequest{}
		require.NoError(t, serviceClient.SetNTPServers(in))
	}
}

func timeHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		time_GetInfo(w)
	case "SetNTPServers":
		time_SetNTPServers(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func time_GetInfo(w http.ResponseWriter) {
	out := time.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func time_SetNTPServers(w http.ResponseWriter) {
	out := time.SetNTPServersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
