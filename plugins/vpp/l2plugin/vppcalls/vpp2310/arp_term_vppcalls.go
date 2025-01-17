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
	"net"

	vpp_l2 "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/l2"
)

func (h *BridgeDomainVppHandler) callBdIPMacAddDel(isAdd bool, bdID uint32, mac string, ip string) error {
	ipAddr, err := ipToAddress(ip)
	if err != nil {
		return err
	}
	macAddr, err := net.ParseMAC(mac)
	if err != nil {
		return err
	}
	bdEntry := vpp_l2.BdIPMac{
		BdID: bdID,
		IP:   ipAddr,
	}
	copy(bdEntry.Mac[:], macAddr)

	req := &vpp_l2.BdIPMacAddDel{
		IsAdd: isAdd,
		Entry: bdEntry,
	}

	reply := &vpp_l2.BdIPMacAddDelReply{}
	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return nil
}

// AddArpTerminationTableEntry creates ARP termination entry for bridge domain.
func (h *BridgeDomainVppHandler) AddArpTerminationTableEntry(bdID uint32, mac string, ip string) error {
	err := h.callBdIPMacAddDel(true, bdID, mac, ip)
	if err != nil {
		return err
	}
	return nil
}

// RemoveArpTerminationTableEntry removes ARP termination entry from bridge domain.
func (h *BridgeDomainVppHandler) RemoveArpTerminationTableEntry(bdID uint32, mac string, ip string) error {
	err := h.callBdIPMacAddDel(false, bdID, mac, ip)
	if err != nil {
		return err
	}
	return nil
}
