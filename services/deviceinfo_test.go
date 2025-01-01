// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/deviceinfo"
	"log"
	"net/http"
	"testing"
)

var deviceinfoMock = &mock.ServiceMock{
	Path:       "/upnp/control/deviceinfo",
	HandleFunc: deviceinfoHandler,
}

func TestDeviceInfo(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", deviceinfoMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server())
	client.Debug = true
	serviceClient := &deviceinfo.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceName: "DeviceInfo",
			ServiceType: "urn:dslforum-org:service:DeviceInfo:1",
			ServiceId:   "urn:DeviceInfo-com:serviceId:DeviceInfo1",
			ServiceUrl:  "/upnp/control/deviceinfo",
		},
	}
	{
		out := &deviceinfo.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &deviceinfo.SetProvisioningCodeRequest{}
		require.NoError(t, serviceClient.SetProvisioningCode(in))
	}
	{
		out := &deviceinfo.GetDeviceLogResponse{}
		require.NoError(t, serviceClient.GetDeviceLog(out))
	}
	{
		out := &deviceinfo.GetSecurityPortResponse{}
		require.NoError(t, serviceClient.GetSecurityPort(out))
	}
	{
		out := &deviceinfo.X_AVM_DE_GetDeviceLogPathResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetDeviceLogPath(out))
	}
}

func deviceinfoHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		deviceinfo_GetInfo(w)
	case "SetProvisioningCode":
		deviceinfo_SetProvisioningCode(w)
	case "GetDeviceLog":
		deviceinfo_GetDeviceLog(w)
	case "GetSecurityPort":
		deviceinfo_GetSecurityPort(w)
	case "X_AVM-DE_GetDeviceLogPath":
		deviceinfo_X_AVM_DE_GetDeviceLogPath(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deviceinfo_GetInfo(w http.ResponseWriter) {
	out := deviceinfo.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceinfo_SetProvisioningCode(w http.ResponseWriter) {
	out := deviceinfo.SetProvisioningCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceinfo_GetDeviceLog(w http.ResponseWriter) {
	out := deviceinfo.GetDeviceLogResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceinfo_GetSecurityPort(w http.ResponseWriter) {
	out := deviceinfo.GetSecurityPortResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func deviceinfo_X_AVM_DE_GetDeviceLogPath(w http.ResponseWriter) {
	out := deviceinfo.X_AVM_DE_GetDeviceLogPathResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
