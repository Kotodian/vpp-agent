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

	vpp_arp "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/arp"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/l3plugin/vppcalls"
	l3 "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/l3"
)

// DumpProxyArpRanges implements proxy arp handler.
func (h *ProxyArpVppHandler) DumpProxyArpRanges() (pArpRngs []*vppcalls.ProxyArpRangesDetails, err error) {
	reqCtx := h.callsChannel.SendMultiRequest(&vpp_arp.ProxyArpDump{})

	for {
		proxyArpDetails := &vpp_arp.ProxyArpDetails{}
		stop, err := reqCtx.ReceiveReply(proxyArpDetails)
		if stop {
			break
		}
		if err != nil {
			h.log.Error(err)
			return nil, err
		}

		pArpRngs = append(pArpRngs, &vppcalls.ProxyArpRangesDetails{
			Range: &l3.ProxyARP_Range{
				FirstIpAddr: net.IP(proxyArpDetails.Proxy.Low[:]).To4().String(),
				LastIpAddr:  net.IP(proxyArpDetails.Proxy.Hi[:]).To4().String(),
				VrfId:       proxyArpDetails.Proxy.TableID,
			},
		})
	}

	return pArpRngs, nil
}

// DumpProxyArpInterfaces implements proxy arp handler.
func (h *ProxyArpVppHandler) DumpProxyArpInterfaces() (pArpIfs []*vppcalls.ProxyArpInterfaceDetails, err error) {
	reqCtx := h.callsChannel.SendMultiRequest(&vpp_arp.ProxyArpIntfcDump{})

	for {
		proxyArpDetails := &vpp_arp.ProxyArpIntfcDetails{}
		stop, err := reqCtx.ReceiveReply(proxyArpDetails)
		if stop {
			break
		}
		if err != nil {
			h.log.Error(err)
			return nil, err
		}

		// Interface
		ifName, _, exists := h.ifIndexes.LookupBySwIfIndex(proxyArpDetails.SwIfIndex)
		if !exists {
			h.log.Warnf("Proxy ARP interface dump: missing name for interface index %d", proxyArpDetails.SwIfIndex)
		}

		// Create entry
		pArpIfs = append(pArpIfs, &vppcalls.ProxyArpInterfaceDetails{
			Interface: &l3.ProxyARP_Interface{
				Name: ifName,
			},
			Meta: &vppcalls.ProxyArpInterfaceMeta{
				SwIfIndex: proxyArpDetails.SwIfIndex,
			},
		})

	}

	return pArpIfs, nil
}
