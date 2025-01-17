//  Copyright (c) 2022 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp2310

import (
	vpp_dhcp "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/dhcp"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/interface_types"
)

func (h *InterfaceVppHandler) handleInterfaceDHCP(ifIdx uint32, hostName string, isAdd bool) error {
	req := &vpp_dhcp.DHCPClientConfig{
		IsAdd: isAdd,
		Client: vpp_dhcp.DHCPClient{
			SwIfIndex:     interface_types.InterfaceIndex(ifIdx),
			Hostname:      hostName,
			WantDHCPEvent: true,
		},
	}
	reply := &vpp_dhcp.DHCPClientConfigReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return nil
}

func (h *InterfaceVppHandler) SetInterfaceAsDHCPClient(ifIdx uint32, hostName string) error {
	return h.handleInterfaceDHCP(ifIdx, hostName, true)
}

func (h *InterfaceVppHandler) UnsetInterfaceAsDHCPClient(ifIdx uint32, hostName string) error {
	return h.handleInterfaceDHCP(ifIdx, hostName, false)
}
