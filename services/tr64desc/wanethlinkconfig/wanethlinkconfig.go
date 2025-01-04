// generated from spec version: 1.0
package wanethlinkconfig

import (
	"encoding/xml"
	"github.com/tdrn-org/go-tr064"
)

type ServiceClient struct {
	TR064Client *tr064.Client
	Service     tr064.ServiceDescriptor
}

type GetEthernetLinkStatusRequest struct {
	XMLName      xml.Name `xml:"u:GetEthernetLinkStatusRequest"`
	XMLNameSpace string   `xml:"xmlns:u,attr"`
}

type GetEthernetLinkStatusResponse struct {
	XMLName               xml.Name `xml:"GetEthernetLinkStatusResponse"`
	NewEthernetLinkStatus string   `xml:"NewEthernetLinkStatus"`
}

func (client *ServiceClient) GetEthernetLinkStatus(out *GetEthernetLinkStatusResponse) error {
	in := &GetEthernetLinkStatusRequest{XMLNameSpace: client.Service.Type()}
	return client.TR064Client.InvokeService(client.Service, "GetEthernetLinkStatus", tr064.NewSOAPRequest(in), tr064.NewSOAPResponse(out))
}
