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
	"testing"

	. "github.com/onsi/gomega"
	"go.ligato.io/cn-infra/v2/logging/logrus"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/aclplugin/vppcalls"
	vpp_acl "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/acl"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/acl_types"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/ethernet_types"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/ip_types"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/memclnt"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/ifaceidx"
)

// Test translation of IP rule into ACL Plugin's format
func TestGetIPRuleMatch(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	icmpV4Rule := ctx.aclHandler.getIPRuleMatches(acl_types.ACLRule{
		DstPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{20, 0, 0, 1}),
			},
			Len: 24,
		},
		SrcPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{10, 0, 0, 1}),
			},
			Len: 24,
		},
		Proto: vppcalls.ICMPv4Proto,
	})
	if icmpV4Rule.GetIcmp() == nil {
		t.Fatal("should have icmp match")
	}

	icmpV6Rule := ctx.aclHandler.getIPRuleMatches(acl_types.ACLRule{
		SrcPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP6,
				Un: ip_types.AddressUnionIP6(ip_types.IP6Address{'d', 'e', 'd', 'd', 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
			},
			Len: 64,
		},
		DstPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP6,
				Un: ip_types.AddressUnionIP6(ip_types.IP6Address{'d', 'e', 'd', 'd', 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
			},
			Len: 32,
		},
		Proto: vppcalls.ICMPv6Proto,
	})
	if icmpV6Rule.GetIcmp() == nil {
		t.Fatal("should have icmpv6 match")
	}

	tcpRule := ctx.aclHandler.getIPRuleMatches(acl_types.ACLRule{
		DstPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{20, 0, 0, 1}),
			},
			Len: 24,
		},
		SrcPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{10, 0, 0, 1}),
			},
			Len: 24,
		},
		Proto: vppcalls.TCPProto,
	})
	if tcpRule.GetTcp() == nil {
		t.Fatal("should have tcp match")
	}

	udpRule := ctx.aclHandler.getIPRuleMatches(acl_types.ACLRule{
		DstPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{20, 0, 0, 1}),
			},
			Len: 24,
		},
		SrcPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{10, 0, 0, 1}),
			},
			Len: 24,
		},
		Proto: vppcalls.UDPProto,
	})
	if udpRule.GetUdp() == nil {
		t.Fatal("should have udp match")
	}
}

// Test translation of MACIP rule into ACL Plugin's format
func TestGetMACIPRuleMatches(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	macipV4Rule := ctx.aclHandler.getMACIPRuleMatches(acl_types.MacipACLRule{
		IsPermit:   1,
		SrcMac:     ethernet_types.MacAddress{2, 'd', 'e', 'a', 'd', 2},
		SrcMacMask: ethernet_types.MacAddress{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		SrcPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP4,
				Un: ip_types.AddressUnionIP4(ip_types.IP4Address{10, 0, 0, 1}),
			},
			Len: 32,
		},
	})
	if macipV4Rule.GetSourceMacAddress() == "" {
		t.Fatal("should have mac match")
	}
	macipV6Rule := ctx.aclHandler.getMACIPRuleMatches(acl_types.MacipACLRule{
		IsPermit:   0,
		SrcMac:     ethernet_types.MacAddress{2, 'd', 'e', 'a', 'd', 2},
		SrcMacMask: ethernet_types.MacAddress{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		SrcPrefix: ip_types.Prefix{
			Address: ip_types.Address{
				Af: ip_types.ADDRESS_IP6,
				Un: ip_types.AddressUnionIP6(ip_types.IP6Address{'d', 'e', 'd', 'd', 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
			},
			Len: 64,
		},
	})
	if macipV6Rule.GetSourceMacAddress() == "" {
		t.Fatal("should have mac match")
	}
}

// Test dumping of IP rules
func TestDumpIPACL(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(
		&vpp_acl.ACLDetails{
			ACLIndex: 0,
			Tag:      "acl1",
			Count:    1,
			R:        []acl_types.ACLRule{{IsPermit: 1}},
		},
		&vpp_acl.ACLDetails{
			ACLIndex: 1,
			Tag:      "acl2",
			Count:    2,
			R:        []acl_types.ACLRule{{IsPermit: 0}, {IsPermit: 2}},
		},
		&vpp_acl.ACLDetails{
			ACLIndex: 2,
			Tag:      "acl3",
			Count:    3,
			R:        []acl_types.ACLRule{{IsPermit: 0}, {IsPermit: 1}, {IsPermit: 2}},
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})
	ctx.MockVpp.MockReply(&vpp_acl.ACLInterfaceListDetails{
		SwIfIndex: 1,
		Count:     2,
		NInput:    1,
		Acls:      []uint32{0, 2},
	})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	ctx.ifIndexes.Put("if0", &ifaceidx.IfaceMetadata{SwIfIndex: 1})

	ifaces, err := ctx.aclHandler.DumpACL()
	Expect(err).To(Succeed())
	Expect(ifaces).To(HaveLen(3))
	//Expect(ifaces[0].Identifier.ACLIndex).To(Equal(uint32(0)))
	//Expect(ifaces[0].vppcalls.ACLDetails.Rules[0].AclAction).To(Equal(uint32(1)))
	//Expect(ifaces[1].Identifier.ACLIndex).To(Equal(uint32(1)))
	//Expect(ifaces[2].Identifier.ACLIndex).To(Equal(uint32(2)))
}

// Test dumping of MACIP rules
func TestDumpMACIPACL(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(
		&vpp_acl.MacipACLDetails{
			ACLIndex: 0,
			Tag:      "acl1",
			Count:    1,
			R:        []acl_types.MacipACLRule{{IsPermit: 1}},
		},
		&vpp_acl.MacipACLDetails{
			ACLIndex: 1,
			Tag:      "acl2",
			Count:    2,
			R:        []acl_types.MacipACLRule{{IsPermit: 0}, {IsPermit: 2}},
		},
		&vpp_acl.MacipACLDetails{
			ACLIndex: 2,
			Tag:      "acl3",
			Count:    3,
			R:        []acl_types.MacipACLRule{{IsPermit: 0}, {IsPermit: 1}, {IsPermit: 2}},
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})
	ctx.MockVpp.MockReply(&vpp_acl.MacipACLInterfaceListDetails{
		SwIfIndex: 1,
		Count:     2,
		Acls:      []uint32{0, 2},
	})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	swIfIndexes := ifaceidx.NewIfaceIndex(logrus.DefaultLogger(), "test")
	swIfIndexes.Put("if0", &ifaceidx.IfaceMetadata{SwIfIndex: 1})

	ifaces, err := ctx.aclHandler.DumpMACIPACL()
	Expect(err).To(Succeed())
	Expect(ifaces).To(HaveLen(3))
	//Expect(ifaces[0].Identifier.ACLIndex).To(Equal(uint32(0)))
	//Expect(ifaces[0].vppcalls.ACLDetails.Rules[0].AclAction).To(Equal(uint32(1)))
	//Expect(ifaces[1].Identifier.ACLIndex).To(Equal(uint32(1)))
	//Expect(ifaces[2].Identifier.ACLIndex).To(Equal(uint32(2)))
}

// Test dumping of interfaces with assigned IP rules
func TestDumpACLInterfaces(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(&vpp_acl.ACLInterfaceListDetails{
		SwIfIndex: 1,
		Count:     2,
		NInput:    1,
		Acls:      []uint32{0, 2},
	})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	ctx.ifIndexes.Put("if0", &ifaceidx.IfaceMetadata{SwIfIndex: 1})

	indexes := []uint32{0, 2}
	ifaces, err := ctx.aclHandler.DumpACLInterfaces(indexes)
	Expect(err).To(Succeed())
	Expect(ifaces).To(HaveLen(2))
	Expect(ifaces[0].Ingress).To(Equal([]string{"if0"}))
	Expect(ifaces[2].Egress).To(Equal([]string{"if0"}))
}

// Test dumping of interfaces with assigned MACIP rules
func TestDumpMACIPACLInterfaces(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(&vpp_acl.MacipACLInterfaceListDetails{
		SwIfIndex: 1,
		Count:     2,
		Acls:      []uint32{0, 1},
	})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	ctx.ifIndexes.Put("if0", &ifaceidx.IfaceMetadata{SwIfIndex: 1})

	indexes := []uint32{0, 1}
	ifaces, err := ctx.aclHandler.DumpMACIPACLInterfaces(indexes)
	Expect(err).To(Succeed())
	Expect(ifaces).To(HaveLen(2))
	Expect(ifaces[0].Ingress).To(Equal([]string{"if0"}))
	Expect(ifaces[0].Egress).To(BeNil())
	Expect(ifaces[1].Ingress).To(Equal([]string{"if0"}))
	Expect(ifaces[1].Egress).To(BeNil())
}

// Test dumping of all configured ACLs with IP-type ruleData
func TestDumpIPAcls(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(&vpp_acl.ACLDetails{
		ACLIndex: 0,
		Count:    1,
		R:        []acl_types.ACLRule{{IsPermit: 1}},
	})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	IPRuleACLs, err := ctx.aclHandler.DumpIPAcls()
	Expect(err).To(Succeed())
	Expect(IPRuleACLs).To(HaveLen(1))
}

// Test dumping of all configured ACLs with MACIP-type ruleData
func TestDumpMacIPAcls(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(&vpp_acl.MacipACLDetails{
		ACLIndex: 0,
		Count:    1,
		R:        []acl_types.MacipACLRule{{IsPermit: 1}},
	})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	MacIPRuleACLs, err := ctx.aclHandler.DumpMacIPAcls()
	Expect(err).To(Succeed())
	Expect(MacIPRuleACLs).To(HaveLen(1))
}

func TestDumpInterfaceIPAcls(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(&vpp_acl.ACLInterfaceListDetails{
		SwIfIndex: 0,
		Count:     2,
		NInput:    1,
		Acls:      []uint32{0, 1},
	})
	ctx.MockVpp.MockReply(&vpp_acl.ACLDetails{
		ACLIndex: 0,
		Count:    1,
		R:        []acl_types.ACLRule{{IsPermit: 1}, {IsPermit: 0}},
	})
	ctx.MockVpp.MockReply(&vpp_acl.ACLDetails{
		ACLIndex: 1,
		Count:    1,
		R:        []acl_types.ACLRule{{IsPermit: 2}, {IsPermit: 0}},
	})

	ACLs, err := ctx.aclHandler.DumpInterfaceACLs(0)
	Expect(err).To(Succeed())
	Expect(ACLs).To(HaveLen(2))
}

func TestDumpInterfaceMACIPAcls(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(&vpp_acl.MacipACLInterfaceListDetails{
		SwIfIndex: 0,
		Count:     2,
		Acls:      []uint32{0, 1},
	})
	ctx.MockVpp.MockReply(&vpp_acl.MacipACLDetails{
		ACLIndex: 0,
		Count:    1,
		R:        []acl_types.MacipACLRule{{IsPermit: 1}, {IsPermit: 0}},
	})
	ctx.MockVpp.MockReply(&vpp_acl.MacipACLDetails{
		ACLIndex: 1,
		Count:    1,
		R:        []acl_types.MacipACLRule{{IsPermit: 2}, {IsPermit: 1}},
	})

	ACLs, err := ctx.aclHandler.DumpInterfaceMACIPACLs(0)
	Expect(err).To(Succeed())
	Expect(ACLs).To(HaveLen(2))
}

func TestDumpInterface(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(&vpp_acl.ACLInterfaceListDetails{
		SwIfIndex: 0,
		Count:     2,
		NInput:    1,
		Acls:      []uint32{0, 1},
	})
	IPacls, err := ctx.aclHandler.DumpInterfaceACLList(0)
	Expect(err).To(BeNil())
	Expect(IPacls.Acls).To(HaveLen(2))

	ctx.MockVpp.MockReply(&vpp_acl.ACLInterfaceListDetails{})
	IPacls, err = ctx.aclHandler.DumpInterfaceACLList(0)
	Expect(err).To(BeNil())
	Expect(IPacls.Acls).To(HaveLen(0))

	ctx.MockVpp.MockReply(&vpp_acl.MacipACLInterfaceListDetails{
		SwIfIndex: 0,
		Count:     2,
		Acls:      []uint32{0, 1},
	})
	MACIPacls, err := ctx.aclHandler.DumpInterfaceMACIPACLList(0)
	Expect(err).To(BeNil())
	Expect(MACIPacls.Acls).To(HaveLen(2))

	ctx.MockVpp.MockReply(&vpp_acl.MacipACLInterfaceListDetails{})
	MACIPacls, err = ctx.aclHandler.DumpInterfaceMACIPACLList(0)
	Expect(err).To(BeNil())
	Expect(MACIPacls.Acls).To(HaveLen(0))
}

func TestDumpInterfaces(t *testing.T) {
	ctx := setupACLTest(t)
	defer ctx.teardownACLTest()

	ctx.MockVpp.MockReply(
		&vpp_acl.ACLInterfaceListDetails{
			SwIfIndex: 0,
			Count:     2,
			NInput:    1,
			Acls:      []uint32{0, 1},
		},
		&vpp_acl.ACLInterfaceListDetails{
			SwIfIndex: 1,
			Count:     1,
			NInput:    1,
			Acls:      []uint32{2},
		},
		&vpp_acl.ACLInterfaceListDetails{
			SwIfIndex: 2,
			Count:     2,
			NInput:    1,
			Acls:      []uint32{3, 4},
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})
	ctx.MockVpp.MockReply(&vpp_acl.MacipACLInterfaceListDetails{
		SwIfIndex: 3,
		Count:     2,
		Acls:      []uint32{6, 7},
	},
		&vpp_acl.MacipACLInterfaceListDetails{
			SwIfIndex: 4,
			Count:     1,
			Acls:      []uint32{5},
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	IPacls, MACIPacls, err := ctx.aclHandler.DumpInterfacesLists()
	Expect(err).To(BeNil())
	Expect(IPacls).To(HaveLen(3))
	Expect(MACIPacls).To(HaveLen(2))
}
