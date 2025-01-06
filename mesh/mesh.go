/*
 * Copyright 2024 Holger de Carne
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mesh

type List struct {
	SchemaVersion string `json:"schema_version"`
	Nodes         []Node `json:"nodes"`
}

func (list *List) Connections() []Connection {
	nodeMap := make(map[string]*Node)
	for nodeIndex, node := range list.Nodes {
		nodeMap[node.Uid] = &list.Nodes[nodeIndex]
	}
	connections := make([]Connection, 0)
	for _, right := range nodeMap {
		for _, rightInterface := range right.NodeInterfaces {
			for _, rightLink := range rightInterface.NodeLinks {
				if rightLink.Node2Uid != right.Uid || !rightLink.IsConnected() {
					continue
				}
				left := nodeMap[rightLink.Node1Uid]
				connection := Connection{
					LeftMeshRole:    left.MeshRole,
					LeftDeviceName:  left.DeviceName,
					RightMeshRole:   right.MeshRole,
					RightDeviceName: right.DeviceName,
					InterfaceName:   rightInterface.Name,
					InterfaceType:   rightInterface.Type,
					MaxDataRateRx:   rightLink.MaxDataRateRx,
					MaxDataRateTx:   rightLink.MaxDataRateTx,
					CurDataRateRx:   rightLink.CurDataRateRx,
					CurDataRateTx:   rightLink.CurDataRateTx,
				}
				connections = append(connections, connection)
			}
		}
	}
	return connections
}

type Connection struct {
	LeftMeshRole    string
	LeftDeviceName  string
	RightMeshRole   string
	RightDeviceName string
	InterfaceName   string
	InterfaceType   string
	MaxDataRateRx   int
	MaxDataRateTx   int
	CurDataRateRx   int
	CurDataRateTx   int
}

type Node struct {
	Uid            string      `json:"uid"`
	DeviceName     string      `json:"device_name"`
	IsMeshed       bool        `json:"is_meshed"`
	MeshRole       string      `json:"mesh_role"`
	NodeInterfaces []Interface `json:"node_interfaces"`
}

func (node *Node) IsMaster() bool {
	return node.IsMeshed && node.MeshRole == "master"
}

func (node *Node) IsSlave() bool {
	return node.IsMeshed && node.MeshRole == "slave"
}

type Interface struct {
	Uid       string `json:"uid"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	NodeLinks []Link `json:"node_links"`
}

type Link struct {
	State             string `json:"state"`
	Node1Uid          string `json:"node_1_uid"`
	Node2Uid          string `json:"node_2_uid"`
	NodeInterface1Uid string `json:"node_interface_1_uid"`
	NodeInterface2Uid string `json:"node_interface_2_uid"`
	MaxDataRateRx     int    `json:"max_data_rate_rx"`
	MaxDataRateTx     int    `json:"max_data_rate_tx"`
	CurDataRateRx     int    `json:"cur_data_rate_rx"`
	CurDataRateTx     int    `json:"cur_data_rate_tx"`
}

func (link *Link) IsConnected() bool {
	return link.State == "CONNECTED"
}
