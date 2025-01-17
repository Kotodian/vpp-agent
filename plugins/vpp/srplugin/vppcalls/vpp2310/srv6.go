// Copyright (c) 2022 Bell Canada, Pantheon Technologies and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package vpp2310 contains wrappers over VPP binary APIs to simplify their usage
package vpp2310

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"go.ligato.io/cn-infra/v2/logging"

	vpp_ifs "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/interface"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/interface_types"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/ip_types"
	vpp_sr "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/sr"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/sr_types"
	vpp2310 "go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/vppcalls/vpp2310"
	ifs "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
	srv6 "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/srv6"
)

// Constants for behavior function hardcoded into VPP (there can be also custom behavior functions implemented as VPP plugins)
// Constants are taken from VPP's vnet/srv6/sr.h (names are modified to Golang from original C form in VPP code)
const (
	BehaviorEnd    = iota + 1 // Behavior of simple endpoint
	BehaviorX                 // Behavior of endpoint with Layer-3 cross-connect
	BehaviorT                 // Behavior of endpoint with specific IPv6 table lookup
	BehaviorDfirst            // Unused. Separator in between regular and D
	BehaviorDX2               // Behavior of endpoint with decapulation and Layer-2 cross-connect (or DX2 with egress VLAN rewrite when VLAN notzero - not supported this variant yet)
	BehaviorDX6               // Behavior of endpoint with decapsulation and IPv6 cross-connect
	BehaviorDX4               // Behavior of endpoint with decapsulation and IPv4 cross-connect
	BehaviorDT6               // Behavior of endpoint with decapsulation and specific IPv6 table lookup
	BehaviorDT4               // Behavior of endpoint with decapsulation and specific IPv4 table lookup
	BehaviorLast              // seems unused, note in VPP: "Must always be the last one"
)

// Constants for steering type
// Constants are taken from VPP's vnet/srv6/sr.h (names are modified to Golang from original C form in VPP code)
const (
	SteerTypeL2   = 2
	SteerTypeIPv4 = 4
	SteerTypeIPv6 = 6
)

// Constants for operation of SR policy modify binary API method
const (
	AddSRList            = iota + 1 // Add SR List to an existing SR policy
	DeleteSRList                    // Delete SR List from an existing SR policy
	ModifyWeightOfSRList            // Modify the weight of an existing SR List
)

// AddLocalSid adds local sid <localSID> into VPP
func (h *SRv6VppHandler) AddLocalSid(localSID *srv6.LocalSID) error {
	return h.addDelLocalSid(false, localSID)
}

// DeleteLocalSid deletes local sid <localSID> in VPP
func (h *SRv6VppHandler) DeleteLocalSid(localSID *srv6.LocalSID) error {
	return h.addDelLocalSid(true, localSID)
}

func (h *SRv6VppHandler) addDelLocalSid(deletion bool, localSID *srv6.LocalSID) error {
	h.log.WithFields(logging.Fields{"localSID": localSID.GetSid(), "delete": deletion, "installationVrfID": h.installationVrfID(localSID), "end function": h.endFunction(localSID)}).
		Debug("Adding/deleting Local SID", localSID.GetSid())
	sidAddr, err := parseIPv6(localSID.GetSid()) // parsing to get some standard sid form
	if err != nil {
		return fmt.Errorf("sid address %s is not IPv6 address: %v", localSID.GetSid(), err) // calls from descriptor are already validated
	}
	if !deletion && localSID.GetEndFunctionAd() != nil {
		return h.addSRProxy(sidAddr, localSID)
	}
	var localsid ip_types.IP6Address
	copy(localsid[:], sidAddr.To16())

	req := &vpp_sr.SrLocalsidAddDel{
		IsDel:    deletion,
		Localsid: localsid,
	}
	req.FibTable = localSID.InstallationVrfId // where to install localsid entry/from where to remove installed localsid entry
	if !deletion {
		if err := h.writeEndFunction(req, localSID); err != nil {
			return err
		}
	}
	reply := &vpp_sr.SrLocalsidAddDelReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("vpp call %q returned: %d", reply.GetMessageName(), reply.Retval)
	}

	h.log.WithFields(logging.Fields{"localSID": sidAddr, "delete": deletion, "installationVrfID": h.installationVrfID(localSID), "end function": h.endFunction(localSID)}).
		Debug("Added/deleted Local SID ", sidAddr)

	return nil
}

// addSRProxy adds local sid with SR-proxy end function (End.AD). This functionality has no binary API in VPP, therefore
// CLI commands are used (VPE binary API that calls VPP's CLI).
func (h *SRv6VppHandler) addSRProxy(sidAddr net.IP, localSID *srv6.LocalSID) error {
	// get VPP-internal names of IN and OUT interfaces
	names, err := h.interfaceNameMapping()
	if err != nil {
		return fmt.Errorf("can't convert interface names from etcd to VPP-internal interface names:%v", err)
	}
	outInterface, found := names[localSID.GetEndFunctionAd().OutgoingInterface]
	if !found {
		return fmt.Errorf("can't find VPP-internal name for interface %v (name in etcd)", localSID.GetEndFunctionAd().OutgoingInterface)
	}
	inInterface, found := names[localSID.GetEndFunctionAd().IncomingInterface]
	if !found {
		return fmt.Errorf("can't find VPP-internal name for interface %v (name in etcd)", localSID.GetEndFunctionAd().IncomingInterface)
	}

	// add SR-proxy using VPP CLI
	var cmd string
	if strings.TrimSpace(localSID.GetEndFunctionAd().L3ServiceAddress) == "" { // L2 service
		cmd = fmt.Sprintf("sr localsid address %v fib-table %v behavior end.ad oif %v iif %v", sidAddr, localSID.InstallationVrfId, outInterface, inInterface)
	} else { // L3 service
		cmd = fmt.Sprintf("sr localsid address %v fib-table %v behavior end.ad nh %v oif %v iif %v", sidAddr, localSID.InstallationVrfId, localSID.GetEndFunctionAd().L3ServiceAddress, outInterface, inInterface)
	}
	data, err := h.RunCli(context.TODO(), cmd)
	if err != nil {
		return err
	}
	if len(strings.TrimSpace(string(data))) > 0 {
		return fmt.Errorf("addition of dynamic segment routing proxy failed by returning nonblank space text in CLI: %v", string(data))
	}
	return nil
}

// interfaceNameMapping dumps from VPP internal names of interfaces and uses them to produce mapping from ligato interface names to vpp internal names.
func (h *SRv6VppHandler) interfaceNameMapping() (map[string]string, error) {
	mapping := make(map[string]string)
	reqCtx := h.callsChannel.SendMultiRequest(&vpp_ifs.SwInterfaceDump{})

	for {
		// get next interface info
		ifDetails := &vpp_ifs.SwInterfaceDetails{}
		stop, err := reqCtx.ReceiveReply(ifDetails)
		if stop {
			break // Break from the loop.
		}
		if err != nil {
			return nil, fmt.Errorf("failed to dump interface: %v", err)
		}

		// extract and compute names
		ligatoName := strings.TrimRight(ifDetails.Tag, "\x00")
		vppInternalName := strings.TrimRight(ifDetails.InterfaceName, "\x00")
		if ifDetails.SupSwIfIndex == uint32(ifDetails.SwIfIndex) && // no subinterface (subinterface are not DPDK)
			guessInterfaceType(strings.TrimRight(ifDetails.InterfaceName, "\x00")) == ifs.Interface_DPDK {
			// fill name for physical interfaces (they are mostly without tag)
			ligatoName = vppInternalName
		}

		mapping[ligatoName] = vppInternalName
	}
	return mapping, nil
}

func (h *SRv6VppHandler) installationVrfID(localSID *srv6.LocalSID) string {
	if localSID != nil {
		return fmt.Sprint(localSID.InstallationVrfId)
	}
	return "<nil>"
}

func (h *SRv6VppHandler) endFunction(localSID *srv6.LocalSID) string {
	switch ef := localSID.GetEndFunction().(type) {
	case *srv6.LocalSID_BaseEndFunction:
		return fmt.Sprintf("End{psp: %v}", ef.BaseEndFunction.Psp)
	case *srv6.LocalSID_EndFunctionX:
		return fmt.Sprintf("X{psp: %v, OutgoingInterface: %v, NextHop: %v}", ef.EndFunctionX.Psp, ef.EndFunctionX.OutgoingInterface, ef.EndFunctionX.NextHop)
	case *srv6.LocalSID_EndFunctionT:
		return fmt.Sprintf("T{psp: %v, vrf: %v}", ef.EndFunctionT.Psp, ef.EndFunctionT.VrfId)
	case *srv6.LocalSID_EndFunctionDx2:
		return fmt.Sprintf("DX2{VlanTag: %v, OutgoingInterface: %v}", ef.EndFunctionDx2.VlanTag, ef.EndFunctionDx2.OutgoingInterface)
	case *srv6.LocalSID_EndFunctionDx4:
		return fmt.Sprintf("DX4{OutgoingInterface: %v, NextHop: %v}", ef.EndFunctionDx4.OutgoingInterface, ef.EndFunctionDx4.NextHop)
	case *srv6.LocalSID_EndFunctionDx6:
		return fmt.Sprintf("DX6{OutgoingInterface: %v, NextHop: %v}", ef.EndFunctionDx6.OutgoingInterface, ef.EndFunctionDx6.NextHop)
	case *srv6.LocalSID_EndFunctionDt4:
		return fmt.Sprintf("DT4{vrf: %v}", ef.EndFunctionDt4.VrfId)
	case *srv6.LocalSID_EndFunctionDt6:
		return fmt.Sprintf("DT6{vrf: %v}", ef.EndFunctionDt6.VrfId)
	case *srv6.LocalSID_EndFunctionAd:
		return fmt.Sprintf("AD{L3ServiceAddress: %v, OutgoingInterface: %v, IncomingInterface: %v}", ef.EndFunctionAd.L3ServiceAddress, ef.EndFunctionAd.OutgoingInterface, ef.EndFunctionAd.IncomingInterface)
	case nil:
		return "<nil>"
	default:
		return "unknown end function"
	}
}

func (h *SRv6VppHandler) writeEndFunction(req *vpp_sr.SrLocalsidAddDel, localSID *srv6.LocalSID) error {
	switch ef := localSID.EndFunction.(type) {
	case *srv6.LocalSID_BaseEndFunction:
		req.Behavior = BehaviorEnd
		req.EndPsp = ef.BaseEndFunction.Psp
	case *srv6.LocalSID_EndFunctionX:
		req.Behavior = BehaviorX
		req.EndPsp = ef.EndFunctionX.Psp
		ifMeta, exists := h.ifIndexes.LookupByName(ef.EndFunctionX.OutgoingInterface)
		if !exists {
			return fmt.Errorf("for interface %v doesn't exist sw index", ef.EndFunctionX.OutgoingInterface)
		}
		req.SwIfIndex = interface_types.InterfaceIndex(ifMeta.SwIfIndex)

		nhIP, err := parseIPv6(ef.EndFunctionX.NextHop) // parses also ipv4 addresses but into ipv6 address form
		if err != nil {
			return err
		}
		nhAddr, err := IPToAddress(nhIP.String())
		if err != nil {
			return err
		}
		req.NhAddr = nhAddr // ipv4 address in ipv6 address form?
	case *srv6.LocalSID_EndFunctionT:
		req.Behavior = BehaviorT
		req.EndPsp = ef.EndFunctionT.Psp
		req.SwIfIndex = interface_types.InterfaceIndex(ef.EndFunctionT.VrfId)
	case *srv6.LocalSID_EndFunctionDx2:
		req.Behavior = BehaviorDX2
		req.VlanIndex = ef.EndFunctionDx2.VlanTag
		ifMeta, exists := h.ifIndexes.LookupByName(ef.EndFunctionDx2.OutgoingInterface)
		if !exists {
			return fmt.Errorf("for interface %v doesn't exist sw index", ef.EndFunctionDx2.OutgoingInterface)
		}
		req.SwIfIndex = interface_types.InterfaceIndex(ifMeta.SwIfIndex)
	case *srv6.LocalSID_EndFunctionDx4:
		req.Behavior = BehaviorDX4
		ifMeta, exists := h.ifIndexes.LookupByName(ef.EndFunctionDx4.OutgoingInterface)
		if !exists {
			return fmt.Errorf("for interface %v doesn't exist sw index", ef.EndFunctionDx4.OutgoingInterface)
		}
		req.SwIfIndex = interface_types.InterfaceIndex(ifMeta.SwIfIndex)
		nhAddr, err := parseIPv6(ef.EndFunctionDx4.NextHop) // parses also IPv4
		if err != nil {
			return err
		}
		nhAddr4 := nhAddr.To4()
		if nhAddr4 == nil {
			return fmt.Errorf("next hop of DX4 end function (%v) is not valid IPv4 address", ef.EndFunctionDx4.NextHop)
		}
		var addr ip_types.IP4Address
		copy(addr[:], nhAddr4)
		req.NhAddr.Af = ip_types.ADDRESS_IP4
		req.NhAddr.Un.SetIP4(addr)
	case *srv6.LocalSID_EndFunctionDx6:
		req.Behavior = BehaviorDX6
		ifMeta, exists := h.ifIndexes.LookupByName(ef.EndFunctionDx6.OutgoingInterface)
		if !exists {
			return fmt.Errorf("for interface %v doesn't exist sw index", ef.EndFunctionDx6.OutgoingInterface)
		}
		req.SwIfIndex = interface_types.InterfaceIndex(ifMeta.SwIfIndex)
		nhAddr6, err := parseIPv6(ef.EndFunctionDx6.NextHop)
		if err != nil {
			return err
		}
		var addr ip_types.IP6Address
		copy(addr[:], nhAddr6)
		req.NhAddr.Af = ip_types.ADDRESS_IP6
		req.NhAddr.Un.SetIP6(addr)
	case *srv6.LocalSID_EndFunctionDt4:
		req.Behavior = BehaviorDT4
		req.SwIfIndex = interface_types.InterfaceIndex(ef.EndFunctionDt4.VrfId)
	case *srv6.LocalSID_EndFunctionDt6:
		req.Behavior = BehaviorDT6
		req.SwIfIndex = interface_types.InterfaceIndex(ef.EndFunctionDt6.VrfId)
	case nil:
		return fmt.Errorf("End function not set. Please configure end function for local SID %v ", localSID.GetSid())
	default:
		return fmt.Errorf("unknown end function (model link type %T)", ef) // EndFunction_AD is handled elsewhere
	}

	return nil
}

// SetEncapsSourceAddress sets for SRv6 in VPP the source address used for encapsulated packet
func (h *SRv6VppHandler) SetEncapsSourceAddress(address string) error {
	h.log.Debugf("Configuring encapsulation source address to address %v", address)
	ipAddress, err := parseIPv6(address)
	if err != nil {
		return err
	}
	var encapSrc ip_types.IP6Address
	copy(encapSrc[:], ipAddress.To16())
	req := &vpp_sr.SrSetEncapSource{
		EncapsSource: encapSrc,
	}
	reply := &vpp_sr.SrSetEncapSourceReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("vpp call %q returned: %d", reply.GetMessageName(), reply.Retval)
	}

	h.log.WithFields(logging.Fields{"Encapsulation source address": address}).
		Debug("Encapsulation source address configured.")

	return nil
}

// AddPolicy adds SRv6 policy <policy> into VPP (including all policy's segment lists).
func (h *SRv6VppHandler) AddPolicy(policy *srv6.Policy) error {
	if err := h.addBasePolicyWithFirstSegmentList(policy); err != nil {
		return fmt.Errorf("can't create Policy with first segment list (Policy: %+v): %v", policy, err)
	}
	if err := h.addOtherSegmentLists(policy); err != nil {
		return fmt.Errorf("can't add all segment lists to created policy %+v: %v", policy, err)
	}
	return nil
}

func (h *SRv6VppHandler) addBasePolicyWithFirstSegmentList(policy *srv6.Policy) error {
	h.log.Debugf("Adding SR policy %+v", policy)
	bindingSid, err := parseIPv6(policy.GetBsid()) // already validated
	if err != nil {
		return fmt.Errorf("binding sid address %s is not IPv6 address: %v", policy.GetBsid(), err) // calls from descriptor are already validated
	}
	if len(policy.SegmentLists) == 0 {
		return fmt.Errorf("policy must have defined at least one segment list (Policy: %+v)", policy) // calls from descriptor are already validated
	}
	sids, err := h.convertPolicySegment(policy.SegmentLists[0])
	if err != nil {
		return err
	}
	var BsidAddr ip_types.IP6Address
	copy(BsidAddr[:], bindingSid.To16())
	// Note: Weight in sr.SrPolicyAdd is leftover from API changes that moved weight into sr.Srv6SidList (it is weight of sid list not of the whole policy)
	req := &vpp_sr.SrPolicyAdd{
		BsidAddr: BsidAddr,
		Sids:     *sids,
		IsEncap:  policy.SrhEncapsulation,
		IsSpray:  policy.SprayBehaviour,
		FibTable: policy.InstallationVrfId,
	}
	reply := &vpp_sr.SrPolicyAddReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("vpp call %q returned: %d", reply.GetMessageName(), reply.Retval)
	}

	h.log.WithFields(logging.Fields{"binding SID": bindingSid, "list of next SIDs": policy.SegmentLists[0].Segments}).
		Debug("base SR policy (policy with just one segment list) added")

	return nil
}

func (h *SRv6VppHandler) addOtherSegmentLists(policy *srv6.Policy) error {
	for _, sl := range policy.SegmentLists[1:] {
		if err := h.AddPolicySegmentList(sl, policy); err != nil {
			return fmt.Errorf("failed to add policy segment %+v: %v", sl, err)
		}
	}
	return nil
}

// DeletePolicy deletes SRv6 policy given by binding SID <bindingSid>
func (h *SRv6VppHandler) DeletePolicy(bindingSid net.IP) error {
	h.log.Debugf("Deleting SR policy with binding SID %v ", bindingSid)
	var BsidAddr ip_types.IP6Address
	copy(BsidAddr[:], bindingSid.To16())
	req := &vpp_sr.SrPolicyDel{
		BsidAddr: BsidAddr, // TODO add ability to define policy also by index (SrPolicyIndex)
	}
	reply := &vpp_sr.SrPolicyDelReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("vpp call %q returned: %d", reply.GetMessageName(), reply.Retval)
	}

	h.log.WithFields(logging.Fields{"binding SID": bindingSid}).Debug("SR policy deleted")

	return nil
}

// AddPolicySegmentList adds segment list <segmentList> to SRv6 policy <policy> in VPP
func (h *SRv6VppHandler) AddPolicySegmentList(segmentList *srv6.Policy_SegmentList, policy *srv6.Policy) error {
	h.log.Debugf("Adding segment %+v to SR policy %+v", segmentList, policy)
	err := h.modPolicy(AddSRList, policy, segmentList, 0)
	if err == nil {
		h.log.WithFields(logging.Fields{"binding SID": policy.Bsid, "list of next SIDs": segmentList.Segments}).
			Debug("SR policy modified(added another segment list)")
	}
	return err
}

// DeletePolicySegmentList removes segment list <segmentList> (with VPP-internal index <segmentVPPIndex>) from SRv6 policy <policy> in VPP
func (h *SRv6VppHandler) DeletePolicySegmentList(segmentList *srv6.Policy_SegmentList, segmentVPPIndex uint32, policy *srv6.Policy) error {
	h.log.Debugf("Removing segment %+v (vpp-internal index %v) from SR policy %+v", segmentList, segmentVPPIndex, policy)
	err := h.modPolicy(DeleteSRList, policy, segmentList, segmentVPPIndex)
	if err == nil {
		h.log.WithFields(logging.Fields{"binding SID": policy.Bsid, "list of next SIDs": segmentList.Segments, "segmentListIndex": segmentVPPIndex}).
			Debug("SR policy modified(removed segment list)")
	}
	return err
}

func (h *SRv6VppHandler) modPolicy(operation sr_types.SrPolicyOp, policy *srv6.Policy, segmentList *srv6.Policy_SegmentList, segmentListIndex uint32) error {
	bindingSid, err := parseIPv6(policy.GetBsid())
	if err != nil {
		return fmt.Errorf("binding sid address %s is not IPv6 address: %v", policy.GetBsid(), err) // calls from descriptor are already validated
	}
	sids, err := h.convertPolicySegment(segmentList)
	if err != nil {
		return err
	}

	var BsidAddr ip_types.IP6Address
	copy(BsidAddr[:], bindingSid.To16())
	// Note: Weight in sr.SrPolicyMod is leftover from API changes that moved weight into sr.Srv6SidList (it is weight of sid list not of the whole policy)
	req := &vpp_sr.SrPolicyMod{
		BsidAddr:  BsidAddr, // TODO add ability to define policy also by index (SrPolicyIndex)
		Operation: operation,
		Sids:      *sids,
		FibTable:  policy.InstallationVrfId,
	}
	if operation == DeleteSRList || operation == ModifyWeightOfSRList {
		req.SlIndex = segmentListIndex
	}

	reply := &vpp_sr.SrPolicyModReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("vpp call %q returned: %d", reply.GetMessageName(), reply.Retval)
	}
	return nil
}

func (h *SRv6VppHandler) convertPolicySegment(segmentList *srv6.Policy_SegmentList) (*vpp_sr.Srv6SidList, error) {
	var segments []ip_types.IP6Address
	for _, sid := range segmentList.Segments {
		// parse to IPv6 address
		parserSid, err := parseIPv6(sid)
		if err != nil {
			return nil, err
		}
		// add sid to segment list
		var ipv6Segment ip_types.IP6Address
		copy(ipv6Segment[:], parserSid)
		segments = append(segments, ipv6Segment)
	}
	sidList := &vpp_sr.Srv6SidList{
		NumSids: uint8(len(segments)),
		Weight:  segmentList.Weight,
	}
	copy(sidList.Sids[:], segments)
	return sidList, nil
}

// RetrievePolicyIndexInfo retrieves index of policy <policy> and its segment lists
func (h *SRv6VppHandler) RetrievePolicyIndexInfo(policy *srv6.Policy) (policyIndex uint32, segmentListIndexes map[*srv6.Policy_SegmentList]uint32, err error) {
	// dump sr policies using VPP CLI
	data, err := h.RunCli(context.TODO(), "sh sr policies")
	if err != nil {
		return ^uint32(0), nil, fmt.Errorf("can't dump index data from VPP: %v", err)
	}

	// do necessary parsing to extract index of segment list
	dumpStr := strings.ToLower(string(data))
	segmentListIndexes = make(map[*srv6.Policy_SegmentList]uint32)

	for _, policyStr := range strings.Split(dumpStr, "-----------") {
		policyHeader := regexp.MustCompile(fmt.Sprintf("\\[(\\d+)\\]\\.-\\s+bsid:\\s*%s", strings.ToLower(strings.TrimSpace(policy.GetBsid()))))
		if policyMatch := policyHeader.FindStringSubmatch(policyStr); policyMatch != nil {
			parsed, err := strconv.ParseUint(policyMatch[1], 10, 32)
			if err != nil {
				return ^uint32(0), nil, fmt.Errorf("can't parse policy index %q (dump: %s)", policyMatch[1], dumpStr)
			}
			policyIndex = uint32(parsed)

			for _, sl := range policy.SegmentLists {
				slRE := regexp.MustCompile(fmt.Sprintf("\\[(\\d+)\\].- < %s,[^:>]*> weight: %d", strings.ToLower(strings.Join(sl.Segments, ", ")), sl.Weight))
				if slMatch := slRE.FindStringSubmatch(policyStr); slMatch != nil {
					parsed, err := strconv.ParseUint(slMatch[1], 10, 32)
					if err != nil {
						return ^uint32(0), nil, fmt.Errorf("can't parse segment policy index %q (dump: %s)", slMatch[1], dumpStr)
					}
					segmentListIndexes[sl] = uint32(parsed)
					continue
				}
				return ^uint32(0), nil, fmt.Errorf("can't find index for segment list %+v (policy bsid %v) in dump %q", sl, policy.GetBsid(), dumpStr)
			}
			return policyIndex, segmentListIndexes, nil
		}
	}
	return ^uint32(0), nil, fmt.Errorf("can't find index for policy with bsid %v in dump %q", policy.GetBsid(), dumpStr)
}

// AddSteering sets in VPP steering into SRv6 policy.
func (h *SRv6VppHandler) AddSteering(steering *srv6.Steering) error {
	return h.addDelSteering(false, steering)
}

// RemoveSteering removes in VPP steering into SRv6 policy.
func (h *SRv6VppHandler) RemoveSteering(steering *srv6.Steering) error {
	return h.addDelSteering(true, steering)
}

func (h *SRv6VppHandler) addDelSteering(delete bool, steering *srv6.Steering) error {
	// defining operation strings for logging
	operationProgressing, operationFinished := "Adding", "Added"
	if delete {
		operationProgressing, operationFinished = "Removing", "Removed"
	}

	// logging info about operation with steering
	switch t := steering.Traffic.(type) {
	case *srv6.Steering_L3Traffic_:
		h.log.Debugf("%v steering for l3 traffic with destination %v to SR policy (binding SID %v, policy index %v)",
			operationProgressing, t.L3Traffic.PrefixAddress, steering.GetPolicyBsid(), steering.GetPolicyIndex())
	case *srv6.Steering_L2Traffic_:
		h.log.Debugf("%v steering for l2 traffic from interface %v to SR policy (binding SID %v, policy index %v)",
			operationProgressing, t.L2Traffic.InterfaceName, steering.GetPolicyBsid(), steering.GetPolicyIndex())
	}

	// converting policy reference
	var policyBSIDAddr ip_types.IP6Address // undefined reference
	var policyIndex = uint32(0)            // undefined reference
	switch ref := steering.PolicyRef.(type) {
	case *srv6.Steering_PolicyBsid:
		bsid, err := parseIPv6(ref.PolicyBsid)
		if err != nil {
			return fmt.Errorf("can't parse binding SID %q to IP address: %v ", ref.PolicyBsid, err)
		}
		copy(policyBSIDAddr[:], bsid.To16())
	case *srv6.Steering_PolicyIndex:
		policyIndex = ref.PolicyIndex
	case nil:
		return fmt.Errorf("policy reference must be provided")
	default:
		return fmt.Errorf("unknown policy reference type (link type %+v)", ref)
	}

	// converting target traffic info
	var prefix ip_types.Prefix
	steerType := sr_types.SrSteer(SteerTypeIPv6)
	tableID := uint32(0)
	intIndex := uint32(0)
	switch t := steering.Traffic.(type) {
	case *srv6.Steering_L3Traffic_:
		ip, ipnet, err := net.ParseCIDR(t.L3Traffic.PrefixAddress)
		if err != nil {
			return fmt.Errorf("can't parse ip prefix %q: %v", t.L3Traffic.PrefixAddress, err)
		}
		if ip.To4() != nil { // IPv4 address
			steerType = SteerTypeIPv4
		}
		tableID = t.L3Traffic.InstallationVrfId
		if ip.To16() != nil {
			prefix.Address.Af = ip_types.ADDRESS_IP6
		}
		prefix.Address, _ = vpp2310.IPToAddress(ip.String())
		maskWidth, _ := ipnet.Mask.Size()
		prefix.Len = uint8(maskWidth)
	case *srv6.Steering_L2Traffic_:
		steerType = SteerTypeL2
		ifMeta, exists := h.ifIndexes.LookupByName(t.L2Traffic.InterfaceName)
		if !exists {
			return fmt.Errorf("for interface %v doesn't exist sw index", t.L2Traffic.InterfaceName)
		}
		intIndex = ifMeta.SwIfIndex
	case nil:
		return fmt.Errorf("traffic type must be provided")
	default:
		return fmt.Errorf("unknown traffic type (link type %+v)", t)
	}
	req := &vpp_sr.SrSteeringAddDel{
		IsDel:         delete,
		TableID:       tableID,
		BsidAddr:      policyBSIDAddr,                           // policy (to which we want to steer routing into) identified by policy binding sid (alternativelly it can be used policy index)
		SrPolicyIndex: policyIndex,                              // policy (to which we want to steer routing into) identified by policy index (alternativelly it can be used policy binding sid)
		TrafficType:   steerType,                                // type of traffic to steer
		Prefix:        prefix,                                   // destination prefix address (L3 traffic type only)
		SwIfIndex:     interface_types.InterfaceIndex(intIndex), // incoming interface (L2 traffic type only)
	}
	reply := &vpp_sr.SrSteeringAddDelReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("vpp call %q returned: %d", reply.GetMessageName(), reply.Retval)
	}

	h.log.WithFields(logging.Fields{
		"steer type":         steerType,
		"L3 prefix":          prefix,
		"L2 interface index": intIndex,
		"policy binding SID": policyBSIDAddr,
		"policy index":       policyIndex,
	}).Debugf("%v steering to SR policy ", operationFinished)

	return nil
}

// guessInterfaceType attempts to guess the correct interface type from its internal name (as given by VPP).
// This is required mainly for those interface types, that do not provide dump binary API,
// such as loopback of af_packet.
func guessInterfaceType(ifName string) ifs.Interface_Type {
	switch {
	case strings.HasPrefix(ifName, "loop"),
		strings.HasPrefix(ifName, "local"):
		return ifs.Interface_SOFTWARE_LOOPBACK
	case strings.HasPrefix(ifName, "memif"):
		return ifs.Interface_MEMIF
	case strings.HasPrefix(ifName, "tap"):
		return ifs.Interface_TAP
	case strings.HasPrefix(ifName, "host"):
		return ifs.Interface_AF_PACKET
	case strings.HasPrefix(ifName, "vxlan"):
		return ifs.Interface_VXLAN_TUNNEL
	case strings.HasPrefix(ifName, "ipsec"):
		return ifs.Interface_IPSEC_TUNNEL
	case strings.HasPrefix(ifName, "vmxnet3"):
		return ifs.Interface_VMXNET3_INTERFACE
	default:
		return ifs.Interface_DPDK
	}
}
