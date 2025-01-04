// generated from spec version: 1.0
package time

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
	XMLName                 xml.Name `xml:"GetInfoResponse"`
	NewNTPServer1           string   `xml:"NewNTPServer1"`
	NewNTPServer2           string   `xml:"NewNTPServer2"`
	NewCurrentLocalTime     string   `xml:"NewCurrentLocalTime"`
	NewLocalTimeZone        string   `xml:"NewLocalTimeZone"`
	NewLocalTimeZoneName    string   `xml:"NewLocalTimeZoneName"`
	NewDaylightSavingsUsed  bool     `xml:"NewDaylightSavingsUsed"`
	NewDaylightSavingsStart string   `xml:"NewDaylightSavingsStart"`
	NewDaylightSavingsEnd   string   `xml:"NewDaylightSavingsEnd"`
}

func (client *ServiceClient) GetInfo(out *GetInfoResponse) error {
	in := &GetInfoRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetInfo", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}

type SetNTPServersRequest struct {
	XMLName       xml.Name `xml:"u:SetNTPServersRequest"`
	XMLNameSpace  string   `xml:"xmlns:u,attr"`
	NewNTPServer1 string   `xml:"NewNTPServer1"`
	NewNTPServer2 string   `xml:"NewNTPServer2"`
}

type SetNTPServersResponse struct {
	XMLName xml.Name `xml:"SetNTPServersResponse"`
}

func (client *ServiceClient) SetNTPServers(in *SetNTPServersRequest) error {
	in.XMLNameSpace = client.Service.Type()
	out := &SetNTPServersResponse{}
	return client.TR064Client.InvokeService(client.Service, "SetNTPServers", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
