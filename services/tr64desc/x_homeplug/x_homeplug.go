// generated from spec version: 1.0
package x_homeplug

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetNumberOfDeviceEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfDeviceEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfDeviceEntriesResponse struct {
	XMLName            xml.Name `xml:"GetNumberOfDeviceEntriesResponse"`
	NewNumberOfEntries uint16   `xml:"NewNumberOfEntries"`
}

func (client *ServiceClient) GetNumberOfDeviceEntries(out *GetNumberOfDeviceEntriesResponse) error {
	in := &GetNumberOfDeviceEntriesRequest{XMLNameSpace: client.Service.Type()}
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfDeviceEntries", soapRequest, soapResponse)
}

type GetGenericDeviceEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetGenericDeviceEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetGenericDeviceEntryResponse struct {
	XMLName             xml.Name `xml:"GetGenericDeviceEntryResponse"`
	NewMACAddress       string   `xml:"NewMACAddress"`
	NewActive           bool     `xml:"NewActive"`
	NewName             string   `xml:"NewName"`
	NewModel            string   `xml:"NewModel"`
	NewUpdateAvailable  bool     `xml:"NewUpdateAvailable"`
	NewUpdateSuccessful string   `xml:"NewUpdateSuccessful"`
}

func (client *ServiceClient) GetGenericDeviceEntry(in *GetGenericDeviceEntryRequest, out *GetGenericDeviceEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetGenericDeviceEntry", soapRequest, soapResponse)
}

type GetSpecificDeviceEntryRequest struct {
	XMLName       xml.Name `xml:"u:GetSpecificDeviceEntryRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
}

type GetSpecificDeviceEntryResponse struct {
	XMLName             xml.Name `xml:"GetSpecificDeviceEntryResponse"`
	NewActive           bool     `xml:"NewActive"`
	NewName             string   `xml:"NewName"`
	NewModel            string   `xml:"NewModel"`
	NewUpdateAvailable  bool     `xml:"NewUpdateAvailable"`
	NewUpdateSuccessful string   `xml:"NewUpdateSuccessful"`
}

func (client *ServiceClient) GetSpecificDeviceEntry(in *GetSpecificDeviceEntryRequest, out *GetSpecificDeviceEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "GetSpecificDeviceEntry", soapRequest, soapResponse)
}

type DeviceDoUpdateRequest struct {
	XMLName       xml.Name `xml:"u:DeviceDoUpdateRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewMACAddress string   `xml:"NewMACAddress"`
}

type DeviceDoUpdateResponse struct {
	XMLName xml.Name `xml:"DeviceDoUpdateResponse"`
}

func (client *ServiceClient) DeviceDoUpdate(in *DeviceDoUpdateRequest) error {
	in.XMLNameSpace = client.Service.Type()
	soapRequest := tr064.NewSOAPRequest(in)
	out := &DeviceDoUpdateResponse{}
	soapResponse := tr064.NewSOAPResponse(out)
	return client.TR064Client.InvokeService(client.Service, "DeviceDoUpdate", soapRequest, soapResponse)
}
