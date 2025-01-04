// generated from spec version: 1.0
package wlanconfig

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type SetEnableRequest struct {
	XMLName      xml.Name `xml:"u:SetEnableRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewEnable    bool     `xml:"NewEnable"`
}

type SetEnableResponse struct {
	XMLName xml.Name `xml:"SetEnableResponse"`
}

func (client *ServiceClient) SetEnable(in *SetEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetEnableResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetEnable", soapRequest, soapResponse)
}

type GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoResponse struct {
	XMLName                         xml.Name `xml:"GetInfoResponse"`
	NewEnable                       bool     `xml:"NewEnable"`
	NewStatus                       string   `xml:"NewStatus"`
	NewMaxBitRate                   string   `xml:"NewMaxBitRate"`
	NewChannel                      uint8    `xml:"NewChannel"`
	NewSSID                         string   `xml:"NewSSID"`
	NewBeaconType                   string   `xml:"NewBeaconType"`
	NewX_AVM_DE_PossibleBeaconTypes string   `xml:"NewX_AVM-DE_PossibleBeaconTypes"`
	NewMACAddressControlEnabled     bool     `xml:"NewMACAddressControlEnabled"`
	NewStandard                     string   `xml:"NewStandard"`
	NewBSSID                        string   `xml:"NewBSSID"`
	NewBasicEncryptionModes         string   `xml:"NewBasicEncryptionModes"`
	NewBasicAuthenticationMode      string   `xml:"NewBasicAuthenticationMode"`
	NewMaxCharsSSID                 uint8    `xml:"NewMaxCharsSSID"`
	NewMinCharsSSID                 uint8    `xml:"NewMinCharsSSID"`
	NewAllowedCharsSSID             string   `xml:"NewAllowedCharsSSID"`
	NewMinCharsPSK                  uint8    `xml:"NewMinCharsPSK"`
	NewMaxCharsPSK                  uint8    `xml:"NewMaxCharsPSK"`
	NewAllowedCharsPSK              string   `xml:"NewAllowedCharsPSK"`
	NewX_AVM_DE_FrequencyBand       string   `xml:"NewX_AVM-DE_FrequencyBand"`
	NewX_AVM_DE_WLANGlobalEnable    bool     `xml:"NewX_AVM-DE_WLANGlobalEnable"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetInfo", soapRequest, soapResponse)
}

type SetConfigRequest struct {
	XMLName                     xml.Name `xml:"u:SetConfigRequest"`
	XMLNameSpace                string   `xml:"xmlns:u,attr"`
	NewMaxBitRate               string   `xml:"NewMaxBitRate"`
	NewChannel                  uint8    `xml:"NewChannel"`
	NewSSID                     string   `xml:"NewSSID"`
	NewBeaconType               string   `xml:"NewBeaconType"`
	NewMACAddressControlEnabled bool     `xml:"NewMACAddressControlEnabled"`
	NewBasicEncryptionModes     string   `xml:"NewBasicEncryptionModes"`
	NewBasicAuthenticationMode  string   `xml:"NewBasicAuthenticationMode"`
}

type SetConfigResponse struct {
	XMLName xml.Name `xml:"SetConfigResponse"`
}

func (client *ServiceClient) SetConfig(in *SetConfigRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetConfigResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetConfig", soapRequest, soapResponse)
}

type SetSecurityKeysRequest struct {
	XMLName          xml.Name `xml:"u:SetSecurityKeysRequest"`
	XMLNameSpace     string   `xml:"xmlns:u,attr"`
	NewWEPKey0       string   `xml:"NewWEPKey0"`
	NewWEPKey1       string   `xml:"NewWEPKey1"`
	NewWEPKey2       string   `xml:"NewWEPKey2"`
	NewWEPKey3       string   `xml:"NewWEPKey3"`
	NewPreSharedKey  string   `xml:"NewPreSharedKey"`
	NewKeyPassphrase string   `xml:"NewKeyPassphrase"`
}

type SetSecurityKeysResponse struct {
	XMLName xml.Name `xml:"SetSecurityKeysResponse"`
}

func (client *ServiceClient) SetSecurityKeys(in *SetSecurityKeysRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetSecurityKeysResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetSecurityKeys", soapRequest, soapResponse)
}

type GetSecurityKeysRequest struct {
	XMLName      xml.Name `xml:"u:GetSecurityKeysRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetSecurityKeysResponse struct {
	XMLName          xml.Name `xml:"GetSecurityKeysResponse"`
	NewWEPKey0       string   `xml:"NewWEPKey0"`
	NewWEPKey1       string   `xml:"NewWEPKey1"`
	NewWEPKey2       string   `xml:"NewWEPKey2"`
	NewWEPKey3       string   `xml:"NewWEPKey3"`
	NewPreSharedKey  string   `xml:"NewPreSharedKey"`
	NewKeyPassphrase string   `xml:"NewKeyPassphrase"`
}

func (client *ServiceClient) GetSecurityKeys(out *GetSecurityKeysResponse) error {
	in := &GetSecurityKeysRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetSecurityKeys", soapRequest, soapResponse)
}

type SetDefaultWEPKeyIndexRequest struct {
	XMLName               xml.Name `xml:"u:SetDefaultWEPKeyIndexRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewDefaultWEPKeyIndex uint8    `xml:"NewDefaultWEPKeyIndex"`
}

type SetDefaultWEPKeyIndexResponse struct {
	XMLName xml.Name `xml:"SetDefaultWEPKeyIndexResponse"`
}

func (client *ServiceClient) SetDefaultWEPKeyIndex(in *SetDefaultWEPKeyIndexRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetDefaultWEPKeyIndexResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetDefaultWEPKeyIndex", soapRequest, soapResponse)
}

type GetDefaultWEPKeyIndexRequest struct {
	XMLName      xml.Name `xml:"u:GetDefaultWEPKeyIndexRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDefaultWEPKeyIndexResponse struct {
	XMLName               xml.Name `xml:"GetDefaultWEPKeyIndexResponse"`
	NewDefaultWEPKeyIndex uint8    `xml:"NewDefaultWEPKeyIndex"`
}

func (client *ServiceClient) GetDefaultWEPKeyIndex(out *GetDefaultWEPKeyIndexResponse) error {
	in := &GetDefaultWEPKeyIndexRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetDefaultWEPKeyIndex", soapRequest, soapResponse)
}

type SetBasBeaconSecurityPropertiesRequest struct {
	XMLName                    xml.Name `xml:"u:SetBasBeaconSecurityPropertiesRequest"`
	XMLNameSpace               string   `xml:"xmlns:u,attr"`
	NewBasicEncryptionModes    string   `xml:"NewBasicEncryptionModes"`
	NewBasicAuthenticationMode string   `xml:"NewBasicAuthenticationMode"`
}

type SetBasBeaconSecurityPropertiesResponse struct {
	XMLName xml.Name `xml:"SetBasBeaconSecurityPropertiesResponse"`
}

func (client *ServiceClient) SetBasBeaconSecurityProperties(in *SetBasBeaconSecurityPropertiesRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetBasBeaconSecurityPropertiesResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetBasBeaconSecurityProperties", soapRequest, soapResponse)
}

type GetBasBeaconSecurityPropertiesRequest struct {
	XMLName      xml.Name `xml:"u:GetBasBeaconSecurityPropertiesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetBasBeaconSecurityPropertiesResponse struct {
	XMLName                    xml.Name `xml:"GetBasBeaconSecurityPropertiesResponse"`
	NewBasicEncryptionModes    string   `xml:"NewBasicEncryptionModes"`
	NewBasicAuthenticationMode string   `xml:"NewBasicAuthenticationMode"`
}

func (client *ServiceClient) GetBasBeaconSecurityProperties(out *GetBasBeaconSecurityPropertiesResponse) error {
	in := &GetBasBeaconSecurityPropertiesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetBasBeaconSecurityProperties", soapRequest, soapResponse)
}

type GetStatisticsRequest struct {
	XMLName      xml.Name `xml:"u:GetStatisticsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetStatisticsResponse struct {
	XMLName                 xml.Name `xml:"GetStatisticsResponse"`
	NewTotalPacketsSent     uint32   `xml:"NewTotalPacketsSent"`
	NewTotalPacketsReceived uint32   `xml:"NewTotalPacketsReceived"`
}

func (client *ServiceClient) GetStatistics(out *GetStatisticsResponse) error {
	in := &GetStatisticsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetStatistics", soapRequest, soapResponse)
}

type GetPacketStatisticsRequest struct {
	XMLName      xml.Name `xml:"u:GetPacketStatisticsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetPacketStatisticsResponse struct {
	XMLName                 xml.Name `xml:"GetPacketStatisticsResponse"`
	NewTotalPacketsSent     uint32   `xml:"NewTotalPacketsSent"`
	NewTotalPacketsReceived uint32   `xml:"NewTotalPacketsReceived"`
}

func (client *ServiceClient) GetPacketStatistics(out *GetPacketStatisticsResponse) error {
	in := &GetPacketStatisticsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetPacketStatistics", soapRequest, soapResponse)
}

type GetBSSIDRequest struct {
	XMLName      xml.Name `xml:"u:GetBSSIDRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetBSSIDResponse struct {
	XMLName  xml.Name `xml:"GetBSSIDResponse"`
	NewBSSID string   `xml:"NewBSSID"`
}

func (client *ServiceClient) GetBSSID(out *GetBSSIDResponse) error {
	in := &GetBSSIDRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetBSSID", soapRequest, soapResponse)
}

type GetSSIDRequest struct {
	XMLName      xml.Name `xml:"u:GetSSIDRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetSSIDResponse struct {
	XMLName xml.Name `xml:"GetSSIDResponse"`
	NewSSID string   `xml:"NewSSID"`
}

func (client *ServiceClient) GetSSID(out *GetSSIDResponse) error {
	in := &GetSSIDRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetSSID", soapRequest, soapResponse)
}

type SetSSIDRequest struct {
	XMLName      xml.Name `xml:"u:SetSSIDRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewSSID      string   `xml:"NewSSID"`
}

type SetSSIDResponse struct {
	XMLName xml.Name `xml:"SetSSIDResponse"`
}

func (client *ServiceClient) SetSSID(in *SetSSIDRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetSSIDResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetSSID", soapRequest, soapResponse)
}

type GetBeaconTypeRequest struct {
	XMLName      xml.Name `xml:"u:GetBeaconTypeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetBeaconTypeResponse struct {
	XMLName                         xml.Name `xml:"GetBeaconTypeResponse"`
	NewBeaconType                   string   `xml:"NewBeaconType"`
	NewX_AVM_DE_PossibleBeaconTypes string   `xml:"NewX_AVM-DE_PossibleBeaconTypes"`
}

func (client *ServiceClient) GetBeaconType(out *GetBeaconTypeResponse) error {
	in := &GetBeaconTypeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetBeaconType", soapRequest, soapResponse)
}

type SetBeaconTypeRequest struct {
	XMLName       xml.Name `xml:"u:SetBeaconTypeRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewBeaconType string   `xml:"NewBeaconType"`
}

type SetBeaconTypeResponse struct {
	XMLName xml.Name `xml:"SetBeaconTypeResponse"`
}

func (client *ServiceClient) SetBeaconType(in *SetBeaconTypeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetBeaconTypeResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetBeaconType", soapRequest, soapResponse)
}

type GetChannelInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetChannelInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetChannelInfoResponse struct {
	XMLName                        xml.Name `xml:"GetChannelInfoResponse"`
	NewChannel                     uint8    `xml:"NewChannel"`
	NewPossibleChannels            string   `xml:"NewPossibleChannels"`
	NewX_AVM_DE_AutoChannelEnabled bool     `xml:"NewX_AVM-DE_AutoChannelEnabled"`
	NewX_AVM_DE_FrequencyBand      string   `xml:"NewX_AVM-DE_FrequencyBand"`
}

func (client *ServiceClient) GetChannelInfo(out *GetChannelInfoResponse) error {
	in := &GetChannelInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetChannelInfo", soapRequest, soapResponse)
}

type SetChannelRequest struct {
	XMLName      xml.Name `xml:"u:SetChannelRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewChannel   uint8    `xml:"NewChannel"`
}

type SetChannelResponse struct {
	XMLName xml.Name `xml:"SetChannelResponse"`
}

func (client *ServiceClient) SetChannel(in *SetChannelRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetChannelResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetChannel", soapRequest, soapResponse)
}

type GetBeaconAdvertisementRequest struct {
	XMLName      xml.Name `xml:"u:GetBeaconAdvertisementRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetBeaconAdvertisementResponse struct {
	XMLName                       xml.Name `xml:"GetBeaconAdvertisementResponse"`
	NewBeaconAdvertisementEnabled bool     `xml:"NewBeaconAdvertisementEnabled"`
}

func (client *ServiceClient) GetBeaconAdvertisement(out *GetBeaconAdvertisementResponse) error {
	in := &GetBeaconAdvertisementRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetBeaconAdvertisement", soapRequest, soapResponse)
}

type SetBeaconAdvertisementRequest struct {
	XMLName                       xml.Name `xml:"u:SetBeaconAdvertisementRequest"`
	XMLNameSpace                  string   `xml:"xmlns:u,attr"`
	NewBeaconAdvertisementEnabled bool     `xml:"NewBeaconAdvertisementEnabled"`
}

type SetBeaconAdvertisementResponse struct {
	XMLName xml.Name `xml:"SetBeaconAdvertisementResponse"`
}

func (client *ServiceClient) SetBeaconAdvertisement(in *SetBeaconAdvertisementRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &SetBeaconAdvertisementResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "SetBeaconAdvertisement", soapRequest, soapResponse)
}

type GetTotalAssociationsRequest struct {
	XMLName      xml.Name `xml:"u:GetTotalAssociationsRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetTotalAssociationsResponse struct {
	XMLName              xml.Name `xml:"GetTotalAssociationsResponse"`
	NewTotalAssociations uint16   `xml:"NewTotalAssociations"`
}

func (client *ServiceClient) GetTotalAssociations(out *GetTotalAssociationsResponse) error {
	in := &GetTotalAssociationsRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetTotalAssociations", soapRequest, soapResponse)
}

type GetGenericAssociatedDeviceInfoRequest struct {
	XMLName                  xml.Name `xml:"u:GetGenericAssociatedDeviceInfoRequest"`
	XMLNameSpace             string   `xml:"xmlns:u,attr"`
	NewAssociatedDeviceIndex uint16   `xml:"NewAssociatedDeviceIndex"`
}

type GetGenericAssociatedDeviceInfoResponse struct {
	XMLName                       xml.Name `xml:"GetGenericAssociatedDeviceInfoResponse"`
	NewAssociatedDeviceMACAddress string   `xml:"NewAssociatedDeviceMACAddress"`
	NewAssociatedDeviceIPAddress  string   `xml:"NewAssociatedDeviceIPAddress"`
	NewAssociatedDeviceAuthState  bool     `xml:"NewAssociatedDeviceAuthState"`
	NewX_AVM_DE_Speed             uint16   `xml:"NewX_AVM-DE_Speed"`
	NewX_AVM_DE_SignalStrength    uint8    `xml:"NewX_AVM-DE_SignalStrength"`
	NewX_AVM_DE_ChannelWidth      uint16   `xml:"NewX_AVM-DE_ChannelWidth"`
}

func (client *ServiceClient) GetGenericAssociatedDeviceInfo(in *GetGenericAssociatedDeviceInfoRequest, out *GetGenericAssociatedDeviceInfoResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetGenericAssociatedDeviceInfo", soapRequest, soapResponse)
}

type GetSpecificAssociatedDeviceInfoRequest struct {
	XMLName                       xml.Name `xml:"u:GetSpecificAssociatedDeviceInfoRequest"`
	XMLNameSpace                  string   `xml:"xmlns:u,attr"`
	NewAssociatedDeviceMACAddress string   `xml:"NewAssociatedDeviceMACAddress"`
}

type GetSpecificAssociatedDeviceInfoResponse struct {
	XMLName                      xml.Name `xml:"GetSpecificAssociatedDeviceInfoResponse"`
	NewAssociatedDeviceIPAddress string   `xml:"NewAssociatedDeviceIPAddress"`
	NewAssociatedDeviceAuthState bool     `xml:"NewAssociatedDeviceAuthState"`
	NewX_AVM_DE_Speed            uint16   `xml:"NewX_AVM-DE_Speed"`
	NewX_AVM_DE_SignalStrength   uint8    `xml:"NewX_AVM-DE_SignalStrength"`
	NewX_AVM_DE_ChannelWidth     uint16   `xml:"NewX_AVM-DE_ChannelWidth"`
}

func (client *ServiceClient) GetSpecificAssociatedDeviceInfo(in *GetSpecificAssociatedDeviceInfoRequest, out *GetSpecificAssociatedDeviceInfoResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetSpecificAssociatedDeviceInfo", soapRequest, soapResponse)
}

type X_AVM_DE_GetSpecificAssociatedDeviceInfoByIpRequest struct {
	XMLName                      xml.Name `xml:"u:X_AVM-DE_GetSpecificAssociatedDeviceInfoByIpRequest"`
	XMLNameSpace                 string   `xml:"xmlns:u,attr"`
	NewAssociatedDeviceIPAddress string   `xml:"NewAssociatedDeviceIPAddress"`
}

type X_AVM_DE_GetSpecificAssociatedDeviceInfoByIpResponse struct {
	XMLName                       xml.Name `xml:"X_AVM-DE_GetSpecificAssociatedDeviceInfoByIpResponse"`
	NewAssociatedDeviceMACAddress string   `xml:"NewAssociatedDeviceMACAddress"`
	NewAssociatedDeviceAuthState  bool     `xml:"NewAssociatedDeviceAuthState"`
	NewX_AVM_DE_Speed             uint16   `xml:"NewX_AVM-DE_Speed"`
	NewX_AVM_DE_SignalStrength    uint8    `xml:"NewX_AVM-DE_SignalStrength"`
	NewX_AVM_DE_ChannelWidth      uint16   `xml:"NewX_AVM-DE_ChannelWidth"`
}

func (client *ServiceClient) X_AVM_DE_GetSpecificAssociatedDeviceInfoByIp(in *X_AVM_DE_GetSpecificAssociatedDeviceInfoByIpRequest, out *X_AVM_DE_GetSpecificAssociatedDeviceInfoByIpResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetSpecificAssociatedDeviceInfoByIp", soapRequest, soapResponse)
}

type X_AVM_DE_GetWLANDeviceListPathRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetWLANDeviceListPathRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetWLANDeviceListPathResponse struct {
	XMLName                        xml.Name `xml:"X_AVM-DE_GetWLANDeviceListPathResponse"`
	NewX_AVM_DE_WLANDeviceListPath string   `xml:"NewX_AVM-DE_WLANDeviceListPath"`
}

func (client *ServiceClient) X_AVM_DE_GetWLANDeviceListPath(out *X_AVM_DE_GetWLANDeviceListPathResponse) error {
	in := &X_AVM_DE_GetWLANDeviceListPathRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetWLANDeviceListPath", soapRequest, soapResponse)
}

type X_AVM_DE_SetStickSurfEnableRequest struct {
	XMLName            xml.Name `xml:"u:X_AVM-DE_SetStickSurfEnableRequest"`
	XMLNameSpace       string   `xml:"xmlns:u,attr"`
	NewStickSurfEnable bool     `xml:"NewStickSurfEnable"`
}

type X_AVM_DE_SetStickSurfEnableResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetStickSurfEnableResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetStickSurfEnable(in *X_AVM_DE_SetStickSurfEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &X_AVM_DE_SetStickSurfEnableResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetStickSurfEnable", soapRequest, soapResponse)
}

type X_AVM_DE_GetIPTVOptimizedRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetIPTVOptimizedRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetIPTVOptimizedResponse struct {
	XMLName                  xml.Name `xml:"X_AVM-DE_GetIPTVOptimizedResponse"`
	NewX_AVM_DE_IPTVoptimize bool     `xml:"NewX_AVM-DE_IPTVoptimize"`
}

func (client *ServiceClient) X_AVM_DE_GetIPTVOptimized(out *X_AVM_DE_GetIPTVOptimizedResponse) error {
	in := &X_AVM_DE_GetIPTVOptimizedRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetIPTVOptimized", soapRequest, soapResponse)
}

type X_AVM_DE_SetIPTVOptimizedRequest struct {
	XMLName                  xml.Name `xml:"u:X_AVM-DE_SetIPTVOptimizedRequest"`
	XMLNameSpace             string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_IPTVoptimize bool     `xml:"NewX_AVM-DE_IPTVoptimize"`
}

type X_AVM_DE_SetIPTVOptimizedResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetIPTVOptimizedResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetIPTVOptimized(in *X_AVM_DE_SetIPTVOptimizedRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &X_AVM_DE_SetIPTVOptimizedResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetIPTVOptimized", soapRequest, soapResponse)
}

type X_AVM_DE_GetNightControlRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetNightControlRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetNightControlResponse struct {
	XMLName                        xml.Name `xml:"X_AVM-DE_GetNightControlResponse"`
	NewNightControl                string   `xml:"NewNightControl"`
	NewNightTimeControlNoForcedOff bool     `xml:"NewNightTimeControlNoForcedOff"`
}

func (client *ServiceClient) X_AVM_DE_GetNightControl(out *X_AVM_DE_GetNightControlResponse) error {
	in := &X_AVM_DE_GetNightControlRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetNightControl", soapRequest, soapResponse)
}

type X_AVM_DE_GetWLANHybridModeRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetWLANHybridModeRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetWLANHybridModeResponse struct {
	XMLName          xml.Name `xml:"X_AVM-DE_GetWLANHybridModeResponse"`
	NewEnable        bool     `xml:"NewEnable"`
	NewBeaconType    string   `xml:"NewBeaconType"`
	NewKeyPassphrase string   `xml:"NewKeyPassphrase"`
	NewSSID          string   `xml:"NewSSID"`
	NewBSSID         string   `xml:"NewBSSID"`
	NewTrafficMode   string   `xml:"NewTrafficMode"`
	NewManualSpeed   bool     `xml:"NewManualSpeed"`
	NewMaxSpeedDS    uint32   `xml:"NewMaxSpeedDS"`
	NewMaxSpeedUS    uint32   `xml:"NewMaxSpeedUS"`
}

func (client *ServiceClient) X_AVM_DE_GetWLANHybridMode(out *X_AVM_DE_GetWLANHybridModeResponse) error {
	in := &X_AVM_DE_GetWLANHybridModeRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetWLANHybridMode", soapRequest, soapResponse)
}

type X_AVM_DE_SetWLANHybridModeRequest struct {
	XMLName          xml.Name `xml:"u:X_AVM-DE_SetWLANHybridModeRequest"`
	XMLNameSpace     string   `xml:"xmlns:u,attr"`
	NewEnable        bool     `xml:"NewEnable"`
	NewBeaconType    string   `xml:"NewBeaconType"`
	NewKeyPassphrase string   `xml:"NewKeyPassphrase"`
	NewSSID          string   `xml:"NewSSID"`
	NewBSSID         string   `xml:"NewBSSID"`
	NewTrafficMode   string   `xml:"NewTrafficMode"`
	NewManualSpeed   bool     `xml:"NewManualSpeed"`
	NewMaxSpeedDS    uint32   `xml:"NewMaxSpeedDS"`
	NewMaxSpeedUS    uint32   `xml:"NewMaxSpeedUS"`
}

type X_AVM_DE_SetWLANHybridModeResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetWLANHybridModeResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetWLANHybridMode(in *X_AVM_DE_SetWLANHybridModeRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &X_AVM_DE_SetWLANHybridModeResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetWLANHybridMode", soapRequest, soapResponse)
}

type X_AVM_DE_GetWLANExtInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetWLANExtInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetWLANExtInfoResponse struct {
	XMLName                      xml.Name `xml:"X_AVM-DE_GetWLANExtInfoResponse"`
	NewX_AVM_DE_APEnabled        string   `xml:"NewX_AVM-DE_APEnabled"`
	NewX_AVM_DE_APType           string   `xml:"NewX_AVM-DE_APType"`
	NewX_AVM_DE_FrequencyBand    string   `xml:"NewX_AVM-DE_FrequencyBand"`
	NewX_AVM_DE_TimeoutActive    string   `xml:"NewX_AVM-DE_TimeoutActive"`
	NewX_AVM_DE_Timeout          string   `xml:"NewX_AVM-DE_Timeout"`
	NewX_AVM_DE_TimeRemain       string   `xml:"NewX_AVM-DE_TimeRemain"`
	NewX_AVM_DE_NoForcedOff      string   `xml:"NewX_AVM-DE_NoForcedOff"`
	NewX_AVM_DE_UserIsolation    string   `xml:"NewX_AVM-DE_UserIsolation"`
	NewX_AVM_DE_EncryptionMode   string   `xml:"NewX_AVM-DE_EncryptionMode"`
	NewX_AVM_DE_LastChangedStamp uint32   `xml:"NewX_AVM-DE_LastChangedStamp"`
}

func (client *ServiceClient) X_AVM_DE_GetWLANExtInfo(out *X_AVM_DE_GetWLANExtInfoResponse) error {
	in := &X_AVM_DE_GetWLANExtInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetWLANExtInfo", soapRequest, soapResponse)
}

type X_AVM_DE_GetWPSInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetWPSInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetWPSInfoResponse struct {
	XMLName               xml.Name `xml:"X_AVM-DE_GetWPSInfoResponse"`
	NewX_AVM_DE_WPSMode   string   `xml:"NewX_AVM-DE_WPSMode"`
	NewX_AVM_DE_WPSStatus string   `xml:"NewX_AVM-DE_WPSStatus"`
}

func (client *ServiceClient) X_AVM_DE_GetWPSInfo(out *X_AVM_DE_GetWPSInfoResponse) error {
	in := &X_AVM_DE_GetWPSInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetWPSInfo", soapRequest, soapResponse)
}

type X_AVM_DE_SetWPSConfigRequest struct {
	XMLName             xml.Name `xml:"u:X_AVM-DE_SetWPSConfigRequest"`
	XMLNameSpace        string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_WPSMode string   `xml:"NewX_AVM-DE_WPSMode"`
}

type X_AVM_DE_SetWPSConfigResponse struct {
	XMLName               xml.Name `xml:"X_AVM-DE_SetWPSConfigResponse"`
	NewX_AVM_DE_WPSStatus string   `xml:"NewX_AVM-DE_WPSStatus"`
}

func (client *ServiceClient) X_AVM_DE_SetWPSConfig(in *X_AVM_DE_SetWPSConfigRequest, out *X_AVM_DE_SetWPSConfigResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetWPSConfig", soapRequest, soapResponse)
}

type X_AVM_DE_SetWPSEnableRequest struct {
	XMLName               xml.Name `xml:"u:X_AVM-DE_SetWPSEnableRequest"`
	XMLNameSpace          string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_WPSEnable bool     `xml:"NewX_AVM-DE_WPSEnable"`
}

type X_AVM_DE_SetWPSEnableResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetWPSEnableResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetWPSEnable(in *X_AVM_DE_SetWPSEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &X_AVM_DE_SetWPSEnableResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetWPSEnable", soapRequest, soapResponse)
}

type X_AVM_DE_SetWLANGlobalEnableRequest struct {
	XMLName                      xml.Name `xml:"u:X_AVM-DE_SetWLANGlobalEnableRequest"`
	XMLNameSpace                 string   `xml:"xmlns:u,attr"`
	NewX_AVM_DE_WLANGlobalEnable bool     `xml:"NewX_AVM-DE_WLANGlobalEnable"`
}

type X_AVM_DE_SetWLANGlobalEnableResponse struct {
	XMLName xml.Name `xml:"X_AVM-DE_SetWLANGlobalEnableResponse"`
}

func (client *ServiceClient) X_AVM_DE_SetWLANGlobalEnable(in *X_AVM_DE_SetWLANGlobalEnableRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &X_AVM_DE_SetWLANGlobalEnableResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_SetWLANGlobalEnable", soapRequest, soapResponse)
}

type X_AVM_DE_GetWLANConnectionInfoRequest struct {
	XMLName      xml.Name `xml:"u:X_AVM-DE_GetWLANConnectionInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type X_AVM_DE_GetWLANConnectionInfoResponse struct {
	XMLName                        xml.Name `xml:"X_AVM-DE_GetWLANConnectionInfoResponse"`
	NewAssociatedDeviceMACAddress  string   `xml:"NewAssociatedDeviceMACAddress"`
	NewSSID                        string   `xml:"NewSSID"`
	NewBSSID                       string   `xml:"NewBSSID"`
	NewBeaconType                  string   `xml:"NewBeaconType"`
	NewChannel                     uint8    `xml:"NewChannel"`
	NewStandard                    string   `xml:"NewStandard"`
	NewX_AVM_DE_AutoChannelEnabled bool     `xml:"NewX_AVM-DE_AutoChannelEnabled"`
	NewX_AVM_DE_ChannelWidth       uint16   `xml:"NewX_AVM-DE_ChannelWidth"`
	NewX_AVM_DE_FrequencyBand      string   `xml:"NewX_AVM-DE_FrequencyBand"`
	NewX_AVM_DE_SignalStrength     uint8    `xml:"NewX_AVM-DE_SignalStrength"`
	NewX_AVM_DE_Speed              uint16   `xml:"NewX_AVM-DE_Speed"`
	NewX_AVM_DE_SpeedRX            uint32   `xml:"NewX_AVM-DE_SpeedRX"`
	NewX_AVM_DE_SpeedMax           uint32   `xml:"NewX_AVM-DE_SpeedMax"`
	NewX_AVM_DE_SpeedRXMax         uint32   `xml:"NewX_AVM-DE_SpeedRXMax"`
}

func (client *ServiceClient) X_AVM_DE_GetWLANConnectionInfo(out *X_AVM_DE_GetWLANConnectionInfoResponse) error {
	in := &X_AVM_DE_GetWLANConnectionInfoRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "X_AVM-DE_GetWLANConnectionInfo", soapRequest, soapResponse)
}
