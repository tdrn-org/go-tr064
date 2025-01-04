// generated from spec version: 1.0
package x_homeauto

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetInfoRequest struct {
	XMLName      xml.Name `xml:"u:GetInfoRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetInfoResponse struct {
	XMLName               xml.Name `xml:"GetInfoResponse"`
	NewAllowedCharsAIN    string   `xml:"NewAllowedCharsAIN"`
	NewMaxCharsAIN        uint16   `xml:"NewMaxCharsAIN"`
	NewMinCharsAIN        uint16   `xml:"NewMinCharsAIN"`
	NewMaxCharsDeviceName uint16   `xml:"NewMaxCharsDeviceName"`
	NewMinCharsDeviceName uint16   `xml:"NewMinCharsDeviceName"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetGenericDeviceInfosRequest struct {
	XMLName      xml.Name `xml:"u:GetGenericDeviceInfosRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetGenericDeviceInfosResponse struct {
	XMLName                   xml.Name `xml:"GetGenericDeviceInfosResponse"`
	NewAIN                    string   `xml:"NewAIN"`
	NewDeviceId               uint16   `xml:"NewDeviceId"`
	NewFunctionBitMask        uint16   `xml:"NewFunctionBitMask"`
	NewFirmwareVersion        string   `xml:"NewFirmwareVersion"`
	NewManufacturer           string   `xml:"NewManufacturer"`
	NewProductName            string   `xml:"NewProductName"`
	NewDeviceName             string   `xml:"NewDeviceName"`
	NewPresent                string   `xml:"NewPresent"`
	NewMultimeterIsEnabled    string   `xml:"NewMultimeterIsEnabled"`
	NewMultimeterIsValid      string   `xml:"NewMultimeterIsValid"`
	NewMultimeterPower        uint32   `xml:"NewMultimeterPower"`
	NewMultimeterEnergy       uint32   `xml:"NewMultimeterEnergy"`
	NewTemperatureIsEnabled   string   `xml:"NewTemperatureIsEnabled"`
	NewTemperatureIsValid     string   `xml:"NewTemperatureIsValid"`
	NewTemperatureCelsius     int32    `xml:"NewTemperatureCelsius"`
	NewTemperatureOffset      int32    `xml:"NewTemperatureOffset"`
	NewSwitchIsEnabled        string   `xml:"NewSwitchIsEnabled"`
	NewSwitchIsValid          string   `xml:"NewSwitchIsValid"`
	NewSwitchState            string   `xml:"NewSwitchState"`
	NewSwitchMode             string   `xml:"NewSwitchMode"`
	NewSwitchLock             bool     `xml:"NewSwitchLock"`
	NewHkrIsEnabled           string   `xml:"NewHkrIsEnabled"`
	NewHkrIsValid             string   `xml:"NewHkrIsValid"`
	NewHkrIsTemperature       int32    `xml:"NewHkrIsTemperature"`
	NewHkrSetVentilStatus     string   `xml:"NewHkrSetVentilStatus"`
	NewHkrSetTemperature      int32    `xml:"NewHkrSetTemperature"`
	NewHkrReduceVentilStatus  string   `xml:"NewHkrReduceVentilStatus"`
	NewHkrReduceTemperature   int32    `xml:"NewHkrReduceTemperature"`
	NewHkrComfortVentilStatus string   `xml:"NewHkrComfortVentilStatus"`
	NewHkrComfortTemperature  int32    `xml:"NewHkrComfortTemperature"`
}

func (client *ServiceClient) GetGenericDeviceInfos(in *GetGenericDeviceInfosRequest, out *GetGenericDeviceInfosResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetGenericDeviceInfos", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetSpecificDeviceInfosRequest struct {
	XMLName      xml.Name `xml:"u:GetSpecificDeviceInfosRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewAIN       string   `xml:"NewAIN"`
}

type GetSpecificDeviceInfosResponse struct {
	XMLName                   xml.Name `xml:"GetSpecificDeviceInfosResponse"`
	NewDeviceId               uint16   `xml:"NewDeviceId"`
	NewFunctionBitMask        uint16   `xml:"NewFunctionBitMask"`
	NewFirmwareVersion        string   `xml:"NewFirmwareVersion"`
	NewManufacturer           string   `xml:"NewManufacturer"`
	NewProductName            string   `xml:"NewProductName"`
	NewDeviceName             string   `xml:"NewDeviceName"`
	NewPresent                string   `xml:"NewPresent"`
	NewMultimeterIsEnabled    string   `xml:"NewMultimeterIsEnabled"`
	NewMultimeterIsValid      string   `xml:"NewMultimeterIsValid"`
	NewMultimeterPower        uint32   `xml:"NewMultimeterPower"`
	NewMultimeterEnergy       uint32   `xml:"NewMultimeterEnergy"`
	NewTemperatureIsEnabled   string   `xml:"NewTemperatureIsEnabled"`
	NewTemperatureIsValid     string   `xml:"NewTemperatureIsValid"`
	NewTemperatureCelsius     int32    `xml:"NewTemperatureCelsius"`
	NewTemperatureOffset      int32    `xml:"NewTemperatureOffset"`
	NewSwitchIsEnabled        string   `xml:"NewSwitchIsEnabled"`
	NewSwitchIsValid          string   `xml:"NewSwitchIsValid"`
	NewSwitchState            string   `xml:"NewSwitchState"`
	NewSwitchMode             string   `xml:"NewSwitchMode"`
	NewSwitchLock             bool     `xml:"NewSwitchLock"`
	NewHkrIsEnabled           string   `xml:"NewHkrIsEnabled"`
	NewHkrIsValid             string   `xml:"NewHkrIsValid"`
	NewHkrIsTemperature       int32    `xml:"NewHkrIsTemperature"`
	NewHkrSetVentilStatus     string   `xml:"NewHkrSetVentilStatus"`
	NewHkrSetTemperature      int32    `xml:"NewHkrSetTemperature"`
	NewHkrReduceVentilStatus  string   `xml:"NewHkrReduceVentilStatus"`
	NewHkrReduceTemperature   int32    `xml:"NewHkrReduceTemperature"`
	NewHkrComfortVentilStatus string   `xml:"NewHkrComfortVentilStatus"`
	NewHkrComfortTemperature  int32    `xml:"NewHkrComfortTemperature"`
}

func (client *ServiceClient) GetSpecificDeviceInfos(in *GetSpecificDeviceInfosRequest, out *GetSpecificDeviceInfosResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetSpecificDeviceInfos", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetDeviceNameRequest struct {
	XMLName       xml.Name `xml:"u:SetDeviceNameRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewAIN        string   `xml:"NewAIN"`
	NewDeviceName string   `xml:"NewDeviceName"`
}

type SetDeviceNameResponse struct {
	XMLName xml.Name `xml:"SetDeviceNameResponse"`
}

func (client *ServiceClient) SetDeviceName(in *SetDeviceNameRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetDeviceNameResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetDeviceName", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetSwitchRequest struct {
	XMLName        xml.Name `xml:"u:SetSwitchRequest"`
	XMLNameSpace   string   `xml:"xmlns:u,attr"`
	NewAIN         string   `xml:"NewAIN"`
	NewSwitchState string   `xml:"NewSwitchState"`
}

type SetSwitchResponse struct {
	XMLName xml.Name `xml:"SetSwitchResponse"`
}

func (client *ServiceClient) SetSwitch(in *SetSwitchRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetSwitchResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetSwitch", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
