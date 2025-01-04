// generated from spec version: 1.0
package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-tr064"
	"github.com/tdrn-org/go-tr064/mock"
	"github.com/tdrn-org/go-tr064/services/tr64desc/wlanconfig"
	"log"
	"net/http"
	"testing"
)

var wlanconfigMock = &mock.ServiceMock{
	Path:       "/upnp/control/wlanconfig1",
	HandleFunc: wlanconfigHandler,
}

func TestWLANConfiguration(t *testing.T) {
	// Start mock server
	tr064Mock := mock.Start("testdata", wlanconfigMock)
	defer tr064Mock.Shutdown()
	// Actual test
	client := tr064.NewClient(tr064Mock.Server(), tr064.ServiceSpec("tr64desc"))
	client.Debug = true
	serviceClient := &wlanconfig.ServiceClient{
		TR064Client: client,
		Service: &tr064.StaticServiceDescriptor{
			ServiceSpec: tr064.ServiceSpec("tr64desc"),
			ServiceId:   "urn:WLANConfiguration-com:serviceId:WLANConfiguration1",
			ServiceUrl:  "/upnp/control/wlanconfig1",
		},
	}
	{
		in := &wlanconfig.SetEnableRequest{}
		require.NoError(t, serviceClient.SetEnable(in))
	}
	{
		out := &wlanconfig.GetInfoResponse{}
		require.NoError(t, serviceClient.GetInfo(out))
	}
	{
		in := &wlanconfig.SetConfigRequest{}
		require.NoError(t, serviceClient.SetConfig(in))
	}
	{
		in := &wlanconfig.SetSecurityKeysRequest{}
		require.NoError(t, serviceClient.SetSecurityKeys(in))
	}
	{
		out := &wlanconfig.GetSecurityKeysResponse{}
		require.NoError(t, serviceClient.GetSecurityKeys(out))
	}
	{
		in := &wlanconfig.SetDefaultWEPKeyIndexRequest{}
		require.NoError(t, serviceClient.SetDefaultWEPKeyIndex(in))
	}
	{
		out := &wlanconfig.GetDefaultWEPKeyIndexResponse{}
		require.NoError(t, serviceClient.GetDefaultWEPKeyIndex(out))
	}
	{
		in := &wlanconfig.SetBasBeaconSecurityPropertiesRequest{}
		require.NoError(t, serviceClient.SetBasBeaconSecurityProperties(in))
	}
	{
		out := &wlanconfig.GetBasBeaconSecurityPropertiesResponse{}
		require.NoError(t, serviceClient.GetBasBeaconSecurityProperties(out))
	}
	{
		out := &wlanconfig.GetStatisticsResponse{}
		require.NoError(t, serviceClient.GetStatistics(out))
	}
	{
		out := &wlanconfig.GetPacketStatisticsResponse{}
		require.NoError(t, serviceClient.GetPacketStatistics(out))
	}
	{
		out := &wlanconfig.GetBSSIDResponse{}
		require.NoError(t, serviceClient.GetBSSID(out))
	}
	{
		out := &wlanconfig.GetSSIDResponse{}
		require.NoError(t, serviceClient.GetSSID(out))
	}
	{
		in := &wlanconfig.SetSSIDRequest{}
		require.NoError(t, serviceClient.SetSSID(in))
	}
	{
		out := &wlanconfig.GetBeaconTypeResponse{}
		require.NoError(t, serviceClient.GetBeaconType(out))
	}
	{
		in := &wlanconfig.SetBeaconTypeRequest{}
		require.NoError(t, serviceClient.SetBeaconType(in))
	}
	{
		out := &wlanconfig.GetChannelInfoResponse{}
		require.NoError(t, serviceClient.GetChannelInfo(out))
	}
	{
		in := &wlanconfig.SetChannelRequest{}
		require.NoError(t, serviceClient.SetChannel(in))
	}
	{
		out := &wlanconfig.GetBeaconAdvertisementResponse{}
		require.NoError(t, serviceClient.GetBeaconAdvertisement(out))
	}
	{
		in := &wlanconfig.SetBeaconAdvertisementRequest{}
		require.NoError(t, serviceClient.SetBeaconAdvertisement(in))
	}
	{
		out := &wlanconfig.GetTotalAssociationsResponse{}
		require.NoError(t, serviceClient.GetTotalAssociations(out))
	}
	{
		in := &wlanconfig.GetGenericAssociatedDeviceInfoRequest{}
		out := &wlanconfig.GetGenericAssociatedDeviceInfoResponse{}
		require.NoError(t, serviceClient.GetGenericAssociatedDeviceInfo(in, out))
	}
	{
		in := &wlanconfig.GetSpecificAssociatedDeviceInfoRequest{}
		out := &wlanconfig.GetSpecificAssociatedDeviceInfoResponse{}
		require.NoError(t, serviceClient.GetSpecificAssociatedDeviceInfo(in, out))
	}
	{
		in := &wlanconfig.X_AVM_DE_GetSpecificAssociatedDeviceInfoByIpRequest{}
		out := &wlanconfig.X_AVM_DE_GetSpecificAssociatedDeviceInfoByIpResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetSpecificAssociatedDeviceInfoByIp(in, out))
	}
	{
		out := &wlanconfig.X_AVM_DE_GetWLANDeviceListPathResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetWLANDeviceListPath(out))
	}
	{
		in := &wlanconfig.X_AVM_DE_SetStickSurfEnableRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetStickSurfEnable(in))
	}
	{
		out := &wlanconfig.X_AVM_DE_GetIPTVOptimizedResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetIPTVOptimized(out))
	}
	{
		in := &wlanconfig.X_AVM_DE_SetIPTVOptimizedRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetIPTVOptimized(in))
	}
	{
		out := &wlanconfig.X_AVM_DE_GetNightControlResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetNightControl(out))
	}
	{
		out := &wlanconfig.X_AVM_DE_GetWLANHybridModeResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetWLANHybridMode(out))
	}
	{
		in := &wlanconfig.X_AVM_DE_SetWLANHybridModeRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetWLANHybridMode(in))
	}
	{
		out := &wlanconfig.X_AVM_DE_GetWLANExtInfoResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetWLANExtInfo(out))
	}
	{
		out := &wlanconfig.X_AVM_DE_GetWPSInfoResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetWPSInfo(out))
	}
	{
		in := &wlanconfig.X_AVM_DE_SetWPSConfigRequest{}
		out := &wlanconfig.X_AVM_DE_SetWPSConfigResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_SetWPSConfig(in, out))
	}
	{
		in := &wlanconfig.X_AVM_DE_SetWPSEnableRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetWPSEnable(in))
	}
	{
		in := &wlanconfig.X_AVM_DE_SetWLANGlobalEnableRequest{}
		require.NoError(t, serviceClient.X_AVM_DE_SetWLANGlobalEnable(in))
	}
	{
		out := &wlanconfig.X_AVM_DE_GetWLANConnectionInfoResponse{}
		require.NoError(t, serviceClient.X_AVM_DE_GetWLANConnectionInfo(out))
	}
}

func wlanconfigHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Mock: ", req.URL)
	action, err := mock.UnmarshalSoapAction(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	switch action {
	case "SetEnable":
		wlanconfig_SetEnable(w)
	case "GetInfo":
		wlanconfig_GetInfo(w)
	case "SetConfig":
		wlanconfig_SetConfig(w)
	case "SetSecurityKeys":
		wlanconfig_SetSecurityKeys(w)
	case "GetSecurityKeys":
		wlanconfig_GetSecurityKeys(w)
	case "SetDefaultWEPKeyIndex":
		wlanconfig_SetDefaultWEPKeyIndex(w)
	case "GetDefaultWEPKeyIndex":
		wlanconfig_GetDefaultWEPKeyIndex(w)
	case "SetBasBeaconSecurityProperties":
		wlanconfig_SetBasBeaconSecurityProperties(w)
	case "GetBasBeaconSecurityProperties":
		wlanconfig_GetBasBeaconSecurityProperties(w)
	case "GetStatistics":
		wlanconfig_GetStatistics(w)
	case "GetPacketStatistics":
		wlanconfig_GetPacketStatistics(w)
	case "GetBSSID":
		wlanconfig_GetBSSID(w)
	case "GetSSID":
		wlanconfig_GetSSID(w)
	case "SetSSID":
		wlanconfig_SetSSID(w)
	case "GetBeaconType":
		wlanconfig_GetBeaconType(w)
	case "SetBeaconType":
		wlanconfig_SetBeaconType(w)
	case "GetChannelInfo":
		wlanconfig_GetChannelInfo(w)
	case "SetChannel":
		wlanconfig_SetChannel(w)
	case "GetBeaconAdvertisement":
		wlanconfig_GetBeaconAdvertisement(w)
	case "SetBeaconAdvertisement":
		wlanconfig_SetBeaconAdvertisement(w)
	case "GetTotalAssociations":
		wlanconfig_GetTotalAssociations(w)
	case "GetGenericAssociatedDeviceInfo":
		wlanconfig_GetGenericAssociatedDeviceInfo(w)
	case "GetSpecificAssociatedDeviceInfo":
		wlanconfig_GetSpecificAssociatedDeviceInfo(w)
	case "X_AVM-DE_GetSpecificAssociatedDeviceInfoByIp":
		wlanconfig_X_AVM_DE_GetSpecificAssociatedDeviceInfoByIp(w)
	case "X_AVM-DE_GetWLANDeviceListPath":
		wlanconfig_X_AVM_DE_GetWLANDeviceListPath(w)
	case "X_AVM-DE_SetStickSurfEnable":
		wlanconfig_X_AVM_DE_SetStickSurfEnable(w)
	case "X_AVM-DE_GetIPTVOptimized":
		wlanconfig_X_AVM_DE_GetIPTVOptimized(w)
	case "X_AVM-DE_SetIPTVOptimized":
		wlanconfig_X_AVM_DE_SetIPTVOptimized(w)
	case "X_AVM-DE_GetNightControl":
		wlanconfig_X_AVM_DE_GetNightControl(w)
	case "X_AVM-DE_GetWLANHybridMode":
		wlanconfig_X_AVM_DE_GetWLANHybridMode(w)
	case "X_AVM-DE_SetWLANHybridMode":
		wlanconfig_X_AVM_DE_SetWLANHybridMode(w)
	case "X_AVM-DE_GetWLANExtInfo":
		wlanconfig_X_AVM_DE_GetWLANExtInfo(w)
	case "X_AVM-DE_GetWPSInfo":
		wlanconfig_X_AVM_DE_GetWPSInfo(w)
	case "X_AVM-DE_SetWPSConfig":
		wlanconfig_X_AVM_DE_SetWPSConfig(w)
	case "X_AVM-DE_SetWPSEnable":
		wlanconfig_X_AVM_DE_SetWPSEnable(w)
	case "X_AVM-DE_SetWLANGlobalEnable":
		wlanconfig_X_AVM_DE_SetWLANGlobalEnable(w)
	case "X_AVM-DE_GetWLANConnectionInfo":
		wlanconfig_X_AVM_DE_GetWLANConnectionInfo(w)

	default:
		log.Println("Unknown action: ", action)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func wlanconfig_SetEnable(w http.ResponseWriter) {
	out := wlanconfig.SetEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetInfo(w http.ResponseWriter) {
	out := wlanconfig.GetInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_SetConfig(w http.ResponseWriter) {
	out := wlanconfig.SetConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_SetSecurityKeys(w http.ResponseWriter) {
	out := wlanconfig.SetSecurityKeysResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetSecurityKeys(w http.ResponseWriter) {
	out := wlanconfig.GetSecurityKeysResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_SetDefaultWEPKeyIndex(w http.ResponseWriter) {
	out := wlanconfig.SetDefaultWEPKeyIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetDefaultWEPKeyIndex(w http.ResponseWriter) {
	out := wlanconfig.GetDefaultWEPKeyIndexResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_SetBasBeaconSecurityProperties(w http.ResponseWriter) {
	out := wlanconfig.SetBasBeaconSecurityPropertiesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetBasBeaconSecurityProperties(w http.ResponseWriter) {
	out := wlanconfig.GetBasBeaconSecurityPropertiesResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetStatistics(w http.ResponseWriter) {
	out := wlanconfig.GetStatisticsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetPacketStatistics(w http.ResponseWriter) {
	out := wlanconfig.GetPacketStatisticsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetBSSID(w http.ResponseWriter) {
	out := wlanconfig.GetBSSIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetSSID(w http.ResponseWriter) {
	out := wlanconfig.GetSSIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_SetSSID(w http.ResponseWriter) {
	out := wlanconfig.SetSSIDResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetBeaconType(w http.ResponseWriter) {
	out := wlanconfig.GetBeaconTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_SetBeaconType(w http.ResponseWriter) {
	out := wlanconfig.SetBeaconTypeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetChannelInfo(w http.ResponseWriter) {
	out := wlanconfig.GetChannelInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_SetChannel(w http.ResponseWriter) {
	out := wlanconfig.SetChannelResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetBeaconAdvertisement(w http.ResponseWriter) {
	out := wlanconfig.GetBeaconAdvertisementResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_SetBeaconAdvertisement(w http.ResponseWriter) {
	out := wlanconfig.SetBeaconAdvertisementResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetTotalAssociations(w http.ResponseWriter) {
	out := wlanconfig.GetTotalAssociationsResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetGenericAssociatedDeviceInfo(w http.ResponseWriter) {
	out := wlanconfig.GetGenericAssociatedDeviceInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_GetSpecificAssociatedDeviceInfo(w http.ResponseWriter) {
	out := wlanconfig.GetSpecificAssociatedDeviceInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_GetSpecificAssociatedDeviceInfoByIp(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_GetSpecificAssociatedDeviceInfoByIpResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_GetWLANDeviceListPath(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_GetWLANDeviceListPathResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_SetStickSurfEnable(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_SetStickSurfEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_GetIPTVOptimized(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_GetIPTVOptimizedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_SetIPTVOptimized(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_SetIPTVOptimizedResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_GetNightControl(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_GetNightControlResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_GetWLANHybridMode(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_GetWLANHybridModeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_SetWLANHybridMode(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_SetWLANHybridModeResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_GetWLANExtInfo(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_GetWLANExtInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_GetWPSInfo(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_GetWPSInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_SetWPSConfig(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_SetWPSConfigResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_SetWPSEnable(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_SetWPSEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_SetWLANGlobalEnable(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_SetWLANGlobalEnableResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}

func wlanconfig_X_AVM_DE_GetWLANConnectionInfo(w http.ResponseWriter) {
	out := wlanconfig.X_AVM_DE_GetWLANConnectionInfoResponse{}
	err := mock.WriteSoapResponse(w, out)
	if err != nil {
		log.Println(err)
	}
}
