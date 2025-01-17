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
	"fmt"
	"net"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/interface_types"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/ip_types"
	vpp_gpe "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/vxlan_gpe"
	ifs "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
)

func (h *InterfaceVppHandler) vxLanGpeAddDelTunnel(isAdd bool, vxLan *ifs.VxlanLink, vrf, multicastIf uint32) (uint32, error) {
	req := &vpp_gpe.VxlanGpeAddDelTunnel{
		McastSwIfIndex: interface_types.InterfaceIndex(multicastIf),
		EncapVrfID:     vrf,
		DecapVrfID:     vxLan.Gpe.DecapVrfId,
		Protocol:       ip_types.IPProto(vxLan.Gpe.Protocol),
		Vni:            vxLan.Vni,
		IsAdd:          isAdd,
	}

	if vxLan.SrcAddress == vxLan.DstAddress {
		return 0, fmt.Errorf("source and destination addresses must not be the same")
	}
	srcAddr := net.ParseIP(vxLan.SrcAddress).To4()
	dstAddr := net.ParseIP(vxLan.DstAddress).To4()
	if srcAddr == nil && dstAddr == nil {
		srcAddr = net.ParseIP(vxLan.SrcAddress).To16()
		dstAddr = net.ParseIP(vxLan.DstAddress).To16()
		if srcAddr == nil || dstAddr == nil {
			return 0, fmt.Errorf("invalid VXLAN address, src: %s, dst: %s", srcAddr, dstAddr)
		}
	} else if srcAddr == nil && dstAddr != nil || srcAddr != nil && dstAddr == nil {
		return 0, fmt.Errorf("IP version mismatch for VXLAN destination and source IP addresses")
	}

	req.Local, _ = IPToAddress(srcAddr.String())
	req.Remote, _ = IPToAddress(dstAddr.String())

	reply := &vpp_gpe.VxlanGpeAddDelTunnelReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return 0, err
	}
	return uint32(reply.SwIfIndex), nil
}

// AddVxLanGpeTunnel adds new VxLAN-GPE interface.
func (h *InterfaceVppHandler) AddVxLanGpeTunnel(ifName string, vrf, multicastIf uint32, vxLan *ifs.VxlanLink) (uint32, error) {
	swIfIndex, err := h.vxLanGpeAddDelTunnel(true, vxLan, vrf, multicastIf)
	if err != nil {
		return 0, err
	}
	return swIfIndex, h.SetInterfaceTag(ifName, swIfIndex)
}

// DeleteVxLanGpeTunnel removes VxLAN-GPE interface.
func (h *InterfaceVppHandler) DeleteVxLanGpeTunnel(ifName string, vxLan *ifs.VxlanLink) error {
	swIfIndex, err := h.vxLanGpeAddDelTunnel(false, vxLan, 0, 0)
	if err != nil {
		return err
	}
	return h.RemoveInterfaceTag(ifName, swIfIndex)
}
