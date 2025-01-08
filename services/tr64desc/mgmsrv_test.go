// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/mgmsrv"
	"log"
	"net/http"
	"testing"
)

var mgmsrvMock = &mock.ServiceMock{
	Path:       "/upnp/control/mgmsrv",
	HandleFunc: mgmsrvHandler,
}

func TestManagementServer(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", mgmsrvMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &mgmsrv.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:ManagementServer:1",
			ServiceId:         "urn:ManagementServer-com:serviceId:ManagementServer1",
			ServiceControlUrl: "/upnp/control/mgmsrv",
		},
	}
	{
		out := &mgmsrv.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &mgmsrv.SetManagementServerURLRequest{}
		require.NoError(t, serviceClient.SetManagementServerURL(in))
	}
	{
		in := &mgmsrv.SetManagementServerUsernameRequest{}
		require.NoError(t, serviceClient.SetManagementServerUsername(in))
	}
	{
		in := &mgmsrv.SetManagementServerPasswordRequest{}
		require.NoError(t, serviceClient.SetManagementServerPassword(in))
	}
	{
		in := &mgmsrv.SetPeriodicInformRequest{}
		require.NoError(t, serviceClient.SetPeriodicInform(in))
	}
	{
		in := &mgmsrv.SetConnectionRequestAuthenticationRequest{}
		require.NoError(t, serviceClient.SetConnectionRequestAuthentication(in))
	}
	{
		in := &mgmsrv.SetUpgradeManagementRequest{}
		require.NoError(t, serviceClient.SetUpgradeManagement(in))
	}
	{
		in := &mgmsrv.X_SetTR069EnableRequest{}
		require.NoError(t, serviceClient.X_SetTR069Enable(in))
	}
	{
		out := &mgmsrv.X_AVM_DE_GetTR069FirmwareDownloadEnabledResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetTR069FirmwareDownloadEnabled(out))
	}
	{
		in := &mgmsrv.X_AVM_DE_SetTR069FirmwareDownloadEnabledRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetTR069FirmwareDownloadEnabled(in))
	}
}

func mgmsrvHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfo":
		mgmsrv_GetInfo(w)
	case "SetManagementServerURL":
		mgmsrv_SetManagementServerURL(w)
	case "SetManagementServerUsername":
		mgmsrv_SetManagementServerUsername(w)
	case "SetManagementServerPassword":
		mgmsrv_SetManagementServerPassword(w)
	case "SetPeriodicInform":
		mgmsrv_SetPeriodicInform(w)
	case "SetConnectionRequestAuthentication":
		mgmsrv_SetConnectionRequestAuthentication(w)
	case "SetUpgradeManagement":
		mgmsrv_SetUpgradeManagement(w)
	case "X_SetTR069Enable":
		mgmsrv_X_SetTR069Enable(w)
	case "X_AVM-DE_GetTR069FirmwareDownloadEnabled":
		mgmsrv_X_AVM_DE_GetTR069FirmwareDownloadEnabled(w)
	case "X_AVM-DE_SetTR069FirmwareDownloadEnabled":
		mgmsrv_X_AVM_DE_SetTR069FirmwareDownloadEnabled(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func mgmsrv_GetInfo(w http.ResponseWriter) {
	out := mgmsrv.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_SetManagementServerURL(w http.ResponseWriter) {
	out := mgmsrv.SetManagementServerURLResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_SetManagementServerUsername(w http.ResponseWriter) {
	out := mgmsrv.SetManagementServerUsernameResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_SetManagementServerPassword(w http.ResponseWriter) {
	out := mgmsrv.SetManagementServerPasswordResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_SetPeriodicInform(w http.ResponseWriter) {
	out := mgmsrv.SetPeriodicInformResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_SetConnectionRequestAuthentication(w http.ResponseWriter) {
	out := mgmsrv.SetConnectionRequestAuthenticationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_SetUpgradeManagement(w http.ResponseWriter) {
	out := mgmsrv.SetUpgradeManagementResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_X_SetTR069Enable(w http.ResponseWriter) {
	out := mgmsrv.X_SetTR069EnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_X_AVM_DE_GetTR069FirmwareDownloadEnabled(w http.ResponseWriter) {
	out := mgmsrv.X_AVM_DE_GetTR069FirmwareDownloadEnabledResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func mgmsrv_X_AVM_DE_SetTR069FirmwareDownloadEnabled(w http.ResponseWriter) {
	out := mgmsrv.X_AVM_DE_SetTR069FirmwareDownloadEnabledResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
