// generated from spec version: 1.0
package x_dect

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetNumberOfDectEntriesRequest struct {
	XMLName      xml.Name `xml:"u:GetNumberOfDectEntriesRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetNumberOfDectEntriesResponse struct {
	XMLName            xml.Name `xml:"GetNumberOfDectEntriesResponse"`
	NewNumberOfEntries uint16   `xml:"NewNumberOfEntries"`
}

func (client *ServiceClient) GetNumberOfDectEntries(out *GetNumberOfDectEntriesResponse) error {
	in := &GetNumberOfDectEntriesRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetNumberOfDectEntries", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetGenericDectEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetGenericDectEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewIndex     uint16   `xml:"NewIndex"`
}

type GetGenericDectEntryResponse struct {
	XMLName             xml.Name `xml:"GetGenericDectEntryResponse"`
	NewID               string   `xml:"NewID"`
	NewActive           bool     `xml:"NewActive"`
	NewName             string   `xml:"NewName"`
	NewModel            string   `xml:"NewModel"`
	NewUpdateAvailable  bool     `xml:"NewUpdateAvailable"`
	NewUpdateSuccessful string   `xml:"NewUpdateSuccessful"`
	NewUpdateInfo       string   `xml:"NewUpdateInfo"`
}

func (client *ServiceClient) GetGenericDectEntry(in *GetGenericDectEntryRequest, out *GetGenericDectEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetGenericDectEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetSpecificDectEntryRequest struct {
	XMLName      xml.Name `xml:"u:GetSpecificDectEntryRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewID        string   `xml:"NewID"`
}

type GetSpecificDectEntryResponse struct {
	XMLName             xml.Name `xml:"GetSpecificDectEntryResponse"`
	NewActive           bool     `xml:"NewActive"`
	NewName             string   `xml:"NewName"`
	NewModel            string   `xml:"NewModel"`
	NewUpdateAvailable  bool     `xml:"NewUpdateAvailable"`
	NewUpdateSuccessful string   `xml:"NewUpdateSuccessful"`
	NewUpdateInfo       string   `xml:"NewUpdateInfo"`
}

func (client *ServiceClient) GetSpecificDectEntry(in *GetSpecificDectEntryRequest, out *GetSpecificDectEntryResponse) error {
	in.XMLNameSpace = client.Service.Type()
	return client.TR064Client.InvokeService(client.Service, "GetSpecificDectEntry", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type DectDoUpdateRequest struct {
	XMLName      xml.Name `xml:"u:DectDoUpdateRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
	NewID        string   `xml:"NewID"`
}

type DectDoUpdateResponse struct {
	XMLName xml.Name `xml:"DectDoUpdateResponse"`
}

func (client *ServiceClient) DectDoUpdate(in *DectDoUpdateRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &DectDoUpdateResponse{}
	return client.TR064Client.InvokeService(client.Service, "DectDoUpdate", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type GetDectListPathRequest struct {
	XMLName      xml.Name `xml:"u:GetDectListPathRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetDectListPathResponse struct {
	XMLName         xml.Name `xml:"GetDectListPathResponse"`
	NewDectListPath string   `xml:"NewDectListPath"`
}

func (client *ServiceClient) GetDectListPath(out *GetDectListPathResponse) error {
	in := &GetDectListPathRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetDectListPath", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
