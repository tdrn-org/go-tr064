// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/x_voip"
	"log"
	"net/http"
	"testing"
)

var x_voipMock = &mock.ServiceMock{
	Path:       "/upnp/control/x_voip",
	HandleFunc: x_voipHandler,
}

func TestX_VoIP(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", x_voipMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &x_voip.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec:       tr064.ServiceSpec("tr64desc"),
			ServiceType:       "urn:dslforum-org:service:X_VoIP:1",
			ServiceId:         "urn:X_VoIP-com:serviceId:X_VoIP1",
			ServiceControlUrl: "/upnp/control/x_voip",
		},
	}
	{
		out := &x_voip.GetInfoExResponse{}
		require.NoError(t, serviceClient.GetInfoEx(out))
	}
	{
		in := &x_voip.X_AVM_DE_AddVoIPAccountRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_AddVoIPAccount(in))
	}
	{
		in := &x_voip.X_AVM_DE_GetVoIPAccountRequest{}
		out := &x_voip.X_AVM_DE_GetVoIPAccountResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetVoIPAccount(in, out))
	}
	{
		in := &x_voip.X_AVM_DE_DelVoIPAccountRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_DelVoIPAccount(in))
	}
	{
		out := &x_voip.X_AVM_DE_GetVoIPAccountsResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetVoIPAccounts(out))
	}
	{
		in := &x_voip.X_AVM_DE_GetVoIPStatusRequest{}
		out := &x_voip.X_AVM_DE_GetVoIPStatusResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetVoIPStatus(in, out))
	}
	{
		out := &x_voip.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &x_voip.SetConfigRequest{}
		require.NoError(t, serviceClient.SetConfig(in))
	}
	{
		out := &x_voip.GetMaxVoIPNumbersResponse{}
		require.NoError(t, serviceClient.GetMaxVoIPNumbers(out))
	}
	{
		out := &x_voip.GetExistingVoIPNumbersResponse{}
		require.NoError(t, serviceClient.GetExistingVoIPNumbers(out))
	}
	{
		out := &x_voip.X_AVM_DE_GetNumberOfClientsResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetNumberOfClients(out))
	}
	{
		in := &x_voip.X_AVM_DE_GetClientRequest{}
		out := &x_voip.X_AVM_DE_GetClientResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetClient(in, out))
	}
	{
		in := &x_voip.X_AVM_DE_GetClient2Request{}
		out := &x_voip.X_AVM_DE_GetClient2Response{}
		require.NoError(t, serviceClient.X_AVM_DE_GetClient2(in, out))
	}
	{
		in := &x_voip.X_AVM_DE_SetClientRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetClient(in))
	}
	{
		in := &x_voip.X_AVM_DE_SetClient2Request{}
		require.NoError(t, serviceClient.X_AVM_DE_SetClient2(in))
	}
	{
		in := &x_voip.X_AVM_DE_GetClient3Request{}
		out := &x_voip.X_AVM_DE_GetClient3Response{}
		require.NoError(t, serviceClient.X_AVM_DE_GetClient3(in, out))
	}
	{
		in := &x_voip.X_AVM_DE_GetClientByClientIdRequest{}
		out := &x_voip.X_AVM_DE_GetClientByClientIdResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetClientByClientId(in, out))
	}
	{
		in := &x_voip.X_AVM_DE_SetClient3Request{}
		require.NoError(t, serviceClient.X_AVM_DE_SetClient3(in))
	}
	{
		in := &x_voip.X_AVM_DE_SetClient4Request{}
		out := &x_voip.X_AVM_DE_SetClient4Response{}
		require.NoError(t, serviceClient.X_AVM_DE_SetClient4(in, out))
	}
	{
		out := &x_voip.X_AVM_DE_GetClientsResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetClients(out))
	}
	{
		out := &x_voip.X_AVM_DE_GetNumberOfNumbersResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetNumberOfNumbers(out))
	}
	{
		out := &x_voip.X_AVM_DE_GetNumbersResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetNumbers(out))
	}
	{
		in := &x_voip.X_AVM_DE_DeleteClientRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_DeleteClient(in))
	}
	{
		out := &x_voip.X_AVM_DE_DialGetConfigResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_DialGetConfig(out))
	}
	{
		in := &x_voip.X_AVM_DE_DialSetConfigRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_DialSetConfig(in))
	}
	{
		in := &x_voip.X_AVM_DE_DialNumberRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_DialNumber(in))
	}
	{
		require.NoError(t, serviceClient.X_AVM_DE_DialHangup())
	}
	{
		in := &x_voip.X_AVM_DE_GetPhonePortRequest{}
		out := &x_voip.X_AVM_DE_GetPhonePortResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetPhonePort(in, out))
	}
	{
		in := &x_voip.X_AVM_DE_SetDelayedCallNotificationRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetDelayedCallNotification(in))
	}
	{
		out := &x_voip.GetVoIPCommonCountryCodeResponse{}
		require.NoError(t, serviceClient.GetVoIPCommonCountryCode(out))
	}
	{
		out := &x_voip.X_AVM_DE_GetVoIPCommonCountryCodeResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetVoIPCommonCountryCode(out))
	}
	{
		in := &x_voip.SetVoIPCommonCountryCodeRequest{}
		require.NoError(t, serviceClient.SetVoIPCommonCountryCode(in))
	}
	{
		in := &x_voip.X_AVM_DE_SetVoIPCommonCountryCodeRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetVoIPCommonCountryCode(in))
	}
	{
		in := &x_voip.GetVoIPEnableCountryCodeRequest{}
		out := &x_voip.GetVoIPEnableCountryCodeResponse{}
		require.NoError(t, serviceClient.GetVoIPEnableCountryCode(in, out))
	}
	{
		in := &x_voip.SetVoIPEnableCountryCodeRequest{}
		require.NoError(t, serviceClient.SetVoIPEnableCountryCode(in))
	}
	{
		out := &x_voip.GetVoIPCommonAreaCodeResponse{}
		require.NoError(t, serviceClient.GetVoIPCommonAreaCode(out))
	}
	{
		out := &x_voip.X_AVM_DE_GetVoIPCommonAreaCodeResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetVoIPCommonAreaCode(out))
	}
	{
		in := &x_voip.SetVoIPCommonAreaCodeRequest{}
		require.NoError(t, serviceClient.SetVoIPCommonAreaCode(in))
	}
	{
		in := &x_voip.X_AVM_DE_SetVoIPCommonAreaCodeRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetVoIPCommonAreaCode(in))
	}
	{
		in := &x_voip.GetVoIPEnableAreaCodeRequest{}
		out := &x_voip.GetVoIPEnableAreaCodeResponse{}
		require.NoError(t, serviceClient.GetVoIPEnableAreaCode(in, out))
	}
	{
		in := &x_voip.SetVoIPEnableAreaCodeRequest{}
		require.NoError(t, serviceClient.SetVoIPEnableAreaCode(in))
	}
	{
		in := &x_voip.X_AVM_DE_GetAlarmClockRequest{}
		out := &x_voip.X_AVM_DE_GetAlarmClockResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetAlarmClock(in, out))
	}
	{
		in := &x_voip.X_AVM_DE_SetAlarmClockEnableRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetAlarmClockEnable(in))
	}
	{
		out := &x_voip.X_AVM_DE_GetNumberOfAlarmClocksResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetNumberOfAlarmClocks(out))
	}
}

func x_voipHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "GetInfoEx":
		x_voip_GetInfoEx(w)
	case "X_AVM-DE_AddVoIPAccount":
		x_voip_X_AVM_DE_AddVoIPAccount(w)
	case "X_AVM-DE_GetVoIPAccount":
		x_voip_X_AVM_DE_GetVoIPAccount(w)
	case "X_AVM-DE_DelVoIPAccount":
		x_voip_X_AVM_DE_DelVoIPAccount(w)
	case "X_AVM-DE_GetVoIPAccounts":
		x_voip_X_AVM_DE_GetVoIPAccounts(w)
	case "X_AVM-DE_GetVoIPStatus":
		x_voip_X_AVM_DE_GetVoIPStatus(w)
	case "GetInfo":
		x_voip_GetInfo(w)
	case "SetConfig":
		x_voip_SetConfig(w)
	case "GetMaxVoIPNumbers":
		x_voip_GetMaxVoIPNumbers(w)
	case "GetExistingVoIPNumbers":
		x_voip_GetExistingVoIPNumbers(w)
	case "X_AVM-DE_GetNumberOfClients":
		x_voip_X_AVM_DE_GetNumberOfClients(w)
	case "X_AVM-DE_GetClient":
		x_voip_X_AVM_DE_GetClient(w)
	case "X_AVM-DE_GetClient2":
		x_voip_X_AVM_DE_GetClient2(w)
	case "X_AVM-DE_SetClient":
		x_voip_X_AVM_DE_SetClient(w)
	case "X_AVM-DE_SetClient2":
		x_voip_X_AVM_DE_SetClient2(w)
	case "X_AVM-DE_GetClient3":
		x_voip_X_AVM_DE_GetClient3(w)
	case "X_AVM-DE_GetClientByClientId":
		x_voip_X_AVM_DE_GetClientByClientId(w)
	case "X_AVM-DE_SetClient3":
		x_voip_X_AVM_DE_SetClient3(w)
	case "X_AVM-DE_SetClient4":
		x_voip_X_AVM_DE_SetClient4(w)
	case "X_AVM-DE_GetClients":
		x_voip_X_AVM_DE_GetClients(w)
	case "X_AVM-DE_GetNumberOfNumbers":
		x_voip_X_AVM_DE_GetNumberOfNumbers(w)
	case "X_AVM-DE_GetNumbers":
		x_voip_X_AVM_DE_GetNumbers(w)
	case "X_AVM-DE_DeleteClient":
		x_voip_X_AVM_DE_DeleteClient(w)
	case "X_AVM-DE_DialGetConfig":
		x_voip_X_AVM_DE_DialGetConfig(w)
	case "X_AVM-DE_DialSetConfig":
		x_voip_X_AVM_DE_DialSetConfig(w)
	case "X_AVM-DE_DialNumber":
		x_voip_X_AVM_DE_DialNumber(w)
	case "X_AVM-DE_DialHangup":
		x_voip_X_AVM_DE_DialHangup(w)
	case "X_AVM-DE_GetPhonePort":
		x_voip_X_AVM_DE_GetPhonePort(w)
	case "X_AVM-DE_SetDelayedCallNotification":
		x_voip_X_AVM_DE_SetDelayedCallNotification(w)
	case "GetVoIPCommonCountryCode":
		x_voip_GetVoIPCommonCountryCode(w)
	case "X_AVM-DE_GetVoIPCommonCountryCode":
		x_voip_X_AVM_DE_GetVoIPCommonCountryCode(w)
	case "SetVoIPCommonCountryCode":
		x_voip_SetVoIPCommonCountryCode(w)
	case "X_AVM-DE_SetVoIPCommonCountryCode":
		x_voip_X_AVM_DE_SetVoIPCommonCountryCode(w)
	case "GetVoIPEnableCountryCode":
		x_voip_GetVoIPEnableCountryCode(w)
	case "SetVoIPEnableCountryCode":
		x_voip_SetVoIPEnableCountryCode(w)
	case "GetVoIPCommonAreaCode":
		x_voip_GetVoIPCommonAreaCode(w)
	case "X_AVM-DE_GetVoIPCommonAreaCode":
		x_voip_X_AVM_DE_GetVoIPCommonAreaCode(w)
	case "SetVoIPCommonAreaCode":
		x_voip_SetVoIPCommonAreaCode(w)
	case "X_AVM-DE_SetVoIPCommonAreaCode":
		x_voip_X_AVM_DE_SetVoIPCommonAreaCode(w)
	case "GetVoIPEnableAreaCode":
		x_voip_GetVoIPEnableAreaCode(w)
	case "SetVoIPEnableAreaCode":
		x_voip_SetVoIPEnableAreaCode(w)
	case "X_AVM-DE_GetAlarmClock":
		x_voip_X_AVM_DE_GetAlarmClock(w)
	case "X_AVM-DE_SetAlarmClockEnable":
		x_voip_X_AVM_DE_SetAlarmClockEnable(w)
	case "X_AVM-DE_GetNumberOfAlarmClocks":
		x_voip_X_AVM_DE_GetNumberOfAlarmClocks(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func x_voip_GetInfoEx(w http.ResponseWriter) {
	out := x_voip.GetInfoExResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_AddVoIPAccount(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_AddVoIPAccountResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetVoIPAccount(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetVoIPAccountResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_DelVoIPAccount(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_DelVoIPAccountResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetVoIPAccounts(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetVoIPAccountsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetVoIPStatus(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetVoIPStatusResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_GetInfo(w http.ResponseWriter) {
	out := x_voip.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_SetConfig(w http.ResponseWriter) {
	out := x_voip.SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_GetMaxVoIPNumbers(w http.ResponseWriter) {
	out := x_voip.GetMaxVoIPNumbersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_GetExistingVoIPNumbers(w http.ResponseWriter) {
	out := x_voip.GetExistingVoIPNumbersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetNumberOfClients(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetNumberOfClientsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetClient(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetClientResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetClient2(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetClient2Response{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_SetClient(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_SetClientResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_SetClient2(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_SetClient2Response{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetClient3(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetClient3Response{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetClientByClientId(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetClientByClientIdResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_SetClient3(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_SetClient3Response{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_SetClient4(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_SetClient4Response{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetClients(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetClientsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetNumberOfNumbers(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetNumberOfNumbersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetNumbers(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetNumbersResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_DeleteClient(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_DeleteClientResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_DialGetConfig(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_DialGetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_DialSetConfig(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_DialSetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_DialNumber(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_DialNumberResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_DialHangup(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_DialHangupResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetPhonePort(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetPhonePortResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_SetDelayedCallNotification(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_SetDelayedCallNotificationResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_GetVoIPCommonCountryCode(w http.ResponseWriter) {
	out := x_voip.GetVoIPCommonCountryCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetVoIPCommonCountryCode(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetVoIPCommonCountryCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_SetVoIPCommonCountryCode(w http.ResponseWriter) {
	out := x_voip.SetVoIPCommonCountryCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_SetVoIPCommonCountryCode(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_SetVoIPCommonCountryCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_GetVoIPEnableCountryCode(w http.ResponseWriter) {
	out := x_voip.GetVoIPEnableCountryCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_SetVoIPEnableCountryCode(w http.ResponseWriter) {
	out := x_voip.SetVoIPEnableCountryCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_GetVoIPCommonAreaCode(w http.ResponseWriter) {
	out := x_voip.GetVoIPCommonAreaCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetVoIPCommonAreaCode(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetVoIPCommonAreaCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_SetVoIPCommonAreaCode(w http.ResponseWriter) {
	out := x_voip.SetVoIPCommonAreaCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_SetVoIPCommonAreaCode(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_SetVoIPCommonAreaCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_GetVoIPEnableAreaCode(w http.ResponseWriter) {
	out := x_voip.GetVoIPEnableAreaCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_SetVoIPEnableAreaCode(w http.ResponseWriter) {
	out := x_voip.SetVoIPEnableAreaCodeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetAlarmClock(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetAlarmClockResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_SetAlarmClockEnable(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_SetAlarmClockEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func x_voip_X_AVM_DE_GetNumberOfAlarmClocks(w http.ResponseWriter) {
	out := x_voip.X_AVM_DE_GetNumberOfAlarmClocksResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
