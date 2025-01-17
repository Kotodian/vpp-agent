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

package vpp2310_test

import (
	"math/bits"
	"net"
	"testing"

	. "github.com/onsi/gomega"

	ifs "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
	nat "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/nat"

	"go.ligato.io/cn-infra/v2/idxmap"
	idxmap_mem "go.ligato.io/cn-infra/v2/idxmap/mem"
	"go.ligato.io/cn-infra/v2/logging/logrus"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/ip_types"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/memclnt"
	vpp_nat_ed "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/nat44_ed"
	vpp_nat_ei "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/nat44_ei"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/nat_types"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/ifaceidx"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/natplugin/vppcalls"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/natplugin/vppcalls/vpp2310"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/vppmock"
)

func TestNat44EdGlobalConfigDump(t *testing.T) {
	ctx, natHandler, swIfIndexes, _ := natTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vpp_nat_ed.Nat44EdPluginEnableDisableReply{})
	err := natHandler.EnableNAT44Plugin(vppcalls.Nat44InitOpts{EndpointDependent: true})
	Expect(err).ShouldNot(HaveOccurred())

	// forwarding
	ctx.MockVpp.MockReply(&vpp_nat_ed.Nat44ShowRunningConfigReply{
		ForwardingEnabled: true,
	})

	// virtual reassembly
	/*ctx.MockVpp.MockReply(&vpp_nat_ed.NatGetReassReply{
		// IPv4
		IP4Timeout:  10,
		IP4MaxReass: 5,
		IP4MaxFrag:  7,
		IP4DropFrag: 1,
		// IPv6
		IP6Timeout:  20,
		IP6MaxReass: 8,
		IP6MaxFrag:  13,
		IP6DropFrag: 0,*
	})*/

	// non-output interfaces
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44InterfaceDetails{
			SwIfIndex: 1,
		},
		&vpp_nat_ed.Nat44InterfaceDetails{
			SwIfIndex: 2,
			Flags:     nat_types.NAT_IS_INSIDE,
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	// output interfaces
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44EdOutputInterfaceDetails{
			SwIfIndex: 3,
		},
		&vpp_nat_ed.Nat44EdOutputInterfaceGetReply{
			Retval: 0,
			Cursor: ^uint32(0),
		})

	// address pool
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44AddressDetails{
			IPAddress: ipTo4Address("192.168.10.1"),
			Flags:     nat_types.NAT_IS_TWICE_NAT,
			VrfID:     1,
		},
		&vpp_nat_ed.Nat44AddressDetails{
			IPAddress: ipTo4Address("192.168.10.2"),
			VrfID:     2,
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	swIfIndexes.Put("if0", &ifaceidx.IfaceMetadata{SwIfIndex: 1})
	swIfIndexes.Put("if1", &ifaceidx.IfaceMetadata{SwIfIndex: 2})
	swIfIndexes.Put("if2", &ifaceidx.IfaceMetadata{SwIfIndex: 3})

	globalCfg, err := natHandler.Nat44GlobalConfigDump(true)
	Expect(err).To(Succeed())

	Expect(globalCfg.Forwarding).To(BeTrue())

	Expect(globalCfg.AddressPool).To(HaveLen(2))
	Expect(globalCfg.AddressPool[0].Address).To(Equal("192.168.10.1"))
	Expect(globalCfg.AddressPool[0].TwiceNat).To(BeTrue())
	Expect(globalCfg.AddressPool[0].VrfId).To(BeEquivalentTo(1))
	Expect(globalCfg.AddressPool[1].Address).To(Equal("192.168.10.2"))
	Expect(globalCfg.AddressPool[1].TwiceNat).To(BeFalse())
	Expect(globalCfg.AddressPool[1].VrfId).To(BeEquivalentTo(2))

	Expect(globalCfg.NatInterfaces).To(HaveLen(3))
	Expect(globalCfg.NatInterfaces[0].Name).To(Equal("if0"))
	Expect(globalCfg.NatInterfaces[0].IsInside).To(BeFalse())
	Expect(globalCfg.NatInterfaces[0].OutputFeature).To(BeFalse())
	Expect(globalCfg.NatInterfaces[1].Name).To(Equal("if1"))
	Expect(globalCfg.NatInterfaces[1].IsInside).To(BeTrue())
	Expect(globalCfg.NatInterfaces[1].OutputFeature).To(BeFalse())
	Expect(globalCfg.NatInterfaces[2].Name).To(Equal("if2"))
	Expect(globalCfg.NatInterfaces[2].IsInside).To(BeFalse())
	Expect(globalCfg.NatInterfaces[2].OutputFeature).To(BeTrue())

	/*Expect(globalCfg.VirtualReassembly).ToNot(BeNil())
	Expect(globalCfg.VirtualReassembly.Timeout).To(BeEquivalentTo(10))
	Expect(globalCfg.VirtualReassembly.MaxReassemblies).To(BeEquivalentTo(5))
	Expect(globalCfg.VirtualReassembly.MaxFragments).To(BeEquivalentTo(7))
	Expect(globalCfg.VirtualReassembly.DropFragments).To(BeTrue())*/
}

func TestNat44EdInterfacesDump(t *testing.T) {
	ctx, natHandler, swIfIndexes, _ := natTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vpp_nat_ed.Nat44EdPluginEnableDisableReply{})
	err := natHandler.EnableNAT44Plugin(vppcalls.Nat44InitOpts{EndpointDependent: true})
	Expect(err).ShouldNot(HaveOccurred())

	// non-output interfaces
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44InterfaceDetails{
			SwIfIndex: 1,
			Flags:     nat_types.NAT_IS_OUTSIDE,
		},
		&vpp_nat_ed.Nat44InterfaceDetails{
			SwIfIndex: 2,
			Flags:     nat_types.NAT_IS_INSIDE,
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	// output interfaces
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44EdOutputInterfaceDetails{
			SwIfIndex: 3,
		},
		&vpp_nat_ed.Nat44EdOutputInterfaceGetReply{
			Retval: 0,
			Cursor: ^uint32(0),
		})

	swIfIndexes.Put("if0", &ifaceidx.IfaceMetadata{SwIfIndex: 1})
	swIfIndexes.Put("if1", &ifaceidx.IfaceMetadata{SwIfIndex: 2})
	swIfIndexes.Put("if2", &ifaceidx.IfaceMetadata{SwIfIndex: 3})

	interfaces, err := natHandler.Nat44InterfacesDump()
	Expect(err).To(Succeed())

	Expect(interfaces).To(HaveLen(3))

	Expect(interfaces[0].Name).To(Equal("if0"))
	Expect(interfaces[0].NatInside).To(BeFalse())
	Expect(interfaces[0].NatOutside).To(BeTrue())
	Expect(interfaces[0].OutputFeature).To(BeFalse())

	Expect(interfaces[1].Name).To(Equal("if1"))
	Expect(interfaces[1].NatInside).To(BeTrue())
	Expect(interfaces[1].NatOutside).To(BeFalse())
	Expect(interfaces[1].OutputFeature).To(BeFalse())

	Expect(interfaces[2].Name).To(Equal("if2"))
	Expect(interfaces[2].NatInside).To(BeFalse())
	Expect(interfaces[2].NatOutside).To(BeFalse())
	Expect(interfaces[2].OutputFeature).To(BeTrue())
}

func TestNat44EiInterfacesDump(t *testing.T) {
	ctx, natHandler, swIfIndexes, _ := natTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vpp_nat_ei.Nat44EiPluginEnableDisableReply{})
	err := natHandler.EnableNAT44Plugin(vppcalls.Nat44InitOpts{EndpointDependent: false})
	Expect(err).ShouldNot(HaveOccurred())

	// non-output interfaces
	ctx.MockVpp.MockReply(
		&vpp_nat_ei.Nat44EiInterfaceDetails{
			SwIfIndex: 1,
			Flags:     vpp_nat_ei.NAT44_EI_IF_OUTSIDE,
		},
		&vpp_nat_ei.Nat44EiInterfaceDetails{
			SwIfIndex: 2,
			Flags:     vpp_nat_ei.NAT44_EI_IF_INSIDE,
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	// output interfaces
	ctx.MockVpp.MockReply(
		&vpp_nat_ei.Nat44EiOutputInterfaceDetails{
			SwIfIndex: 3,
		},
		&vpp_nat_ei.Nat44EiOutputInterfaceGetReply{
			Retval: 0,
			Cursor: ^uint32(0),
		})

	swIfIndexes.Put("if0", &ifaceidx.IfaceMetadata{SwIfIndex: 1})
	swIfIndexes.Put("if1", &ifaceidx.IfaceMetadata{SwIfIndex: 2})
	swIfIndexes.Put("if2", &ifaceidx.IfaceMetadata{SwIfIndex: 3})

	interfaces, err := natHandler.Nat44InterfacesDump()
	Expect(err).To(Succeed())

	Expect(interfaces).To(HaveLen(3))

	Expect(interfaces[0].Name).To(Equal("if0"))
	Expect(interfaces[0].NatInside).To(BeFalse())
	Expect(interfaces[0].NatOutside).To(BeTrue())
	Expect(interfaces[0].OutputFeature).To(BeFalse())

	Expect(interfaces[1].Name).To(Equal("if1"))
	Expect(interfaces[1].NatInside).To(BeTrue())
	Expect(interfaces[1].NatOutside).To(BeFalse())
	Expect(interfaces[1].OutputFeature).To(BeFalse())

	Expect(interfaces[2].Name).To(Equal("if2"))
	Expect(interfaces[2].NatInside).To(BeFalse())
	Expect(interfaces[2].NatOutside).To(BeFalse())
	Expect(interfaces[2].OutputFeature).To(BeTrue())
}

func TestNat44EdAddressPoolsDump(t *testing.T) {
	ctx, natHandler, _, _ := natTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vpp_nat_ed.Nat44EdPluginEnableDisableReply{})
	err := natHandler.EnableNAT44Plugin(vppcalls.Nat44InitOpts{EndpointDependent: true})
	Expect(err).ShouldNot(HaveOccurred())

	// address pool
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44AddressDetails{
			IPAddress: ipTo4Address("192.168.10.1"),
			Flags:     nat_types.NAT_IS_TWICE_NAT,
			VrfID:     1,
		},
		&vpp_nat_ed.Nat44AddressDetails{
			IPAddress: ipTo4Address("192.168.10.2"),
			VrfID:     2,
		},
		&vpp_nat_ed.Nat44AddressDetails{
			IPAddress: ipTo4Address("192.168.10.3"),
			VrfID:     2,
		},
		&vpp_nat_ed.Nat44AddressDetails{
			IPAddress: ipTo4Address("192.168.10.4"),
			VrfID:     2,
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	pools, err := natHandler.Nat44AddressPoolsDump()
	Expect(err).To(Succeed())

	Expect(pools).To(HaveLen(2))

	Expect(pools[0].FirstIp).To(Equal("192.168.10.1"))
	Expect(pools[0].LastIp).To(Equal(""))
	Expect(pools[0].TwiceNat).To(BeTrue())
	Expect(pools[0].VrfId).To(BeEquivalentTo(1))

	Expect(pools[1].FirstIp).To(Equal("192.168.10.2"))
	Expect(pools[1].LastIp).To(Equal("192.168.10.4"))
	Expect(pools[1].TwiceNat).To(BeFalse())
	Expect(pools[1].VrfId).To(BeEquivalentTo(2))
}

func TestNat44EiAddressPoolsDump(t *testing.T) {
	ctx, natHandler, _, _ := natTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vpp_nat_ei.Nat44EiPluginEnableDisableReply{})
	err := natHandler.EnableNAT44Plugin(vppcalls.Nat44InitOpts{EndpointDependent: false})
	Expect(err).ShouldNot(HaveOccurred())

	// address pool
	ctx.MockVpp.MockReply(
		&vpp_nat_ei.Nat44EiAddressDetails{
			IPAddress: ipTo4Address("192.168.10.1"),
			VrfID:     1,
		},
		&vpp_nat_ei.Nat44EiAddressDetails{
			IPAddress: ipTo4Address("192.168.10.2"),
			VrfID:     2,
		},
		&vpp_nat_ei.Nat44EiAddressDetails{
			IPAddress: ipTo4Address("192.168.10.3"),
			VrfID:     2,
		},
		&vpp_nat_ei.Nat44EiAddressDetails{
			IPAddress: ipTo4Address("192.168.10.4"),
			VrfID:     2,
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	pools, err := natHandler.Nat44AddressPoolsDump()
	Expect(err).To(Succeed())

	Expect(pools).To(HaveLen(2))

	Expect(pools[0].FirstIp).To(Equal("192.168.10.1"))
	Expect(pools[0].LastIp).To(Equal(""))
	Expect(pools[0].VrfId).To(BeEquivalentTo(1))

	Expect(pools[1].FirstIp).To(Equal("192.168.10.2"))
	Expect(pools[1].LastIp).To(Equal("192.168.10.4"))
	Expect(pools[1].TwiceNat).To(BeFalse())
	Expect(pools[1].VrfId).To(BeEquivalentTo(2))
}

func TestDNATDump(t *testing.T) {
	ctx, natHandler, swIfIndexes, dhcpIndexes := natTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vpp_nat_ed.Nat44EdPluginEnableDisableReply{})
	err := natHandler.EnableNAT44Plugin(vppcalls.Nat44InitOpts{EndpointDependent: true})
	Expect(err).ShouldNot(HaveOccurred())

	// non-LB static mappings
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44StaticMappingDetails{
			LocalIPAddress:    ipTo4Address("10.10.11.120"),
			ExternalIPAddress: ipTo4Address("10.36.20.20"),
			Protocol:          6,
			LocalPort:         8080,
			ExternalPort:      80,
			ExternalSwIfIndex: vpp2310.NoInterface,
			VrfID:             1,
			Tag:               "DNAT 1",
			Flags:             nat_types.NAT_IS_TWICE_NAT,
		},
		&vpp_nat_ed.Nat44StaticMappingDetails{
			LocalIPAddress:    ipTo4Address("10.10.11.120"),
			Protocol:          6,
			LocalPort:         8080,
			ExternalPort:      80,
			ExternalSwIfIndex: 1,
			VrfID:             1,
			Tag:               "DNAT 1",
			Flags:             nat_types.NAT_IS_TWICE_NAT,
		},
		&vpp_nat_ed.Nat44StaticMappingDetails{
			LocalIPAddress:    ipTo4Address("10.10.11.140"),
			Protocol:          6,
			LocalPort:         8081,
			ExternalPort:      80,
			ExternalSwIfIndex: 2,
			VrfID:             1,
			Tag:               "DNAT 2",
			Flags:             nat_types.NAT_IS_SELF_TWICE_NAT,
		},
		// auto-generated mappings with interface replaced by all assigned IP addresses
		&vpp_nat_ed.Nat44StaticMappingDetails{
			LocalIPAddress:    ipTo4Address("10.10.11.120"),
			ExternalIPAddress: ipTo4Address("10.36.20.30"),
			Protocol:          6,
			LocalPort:         8080,
			ExternalPort:      80,
			ExternalSwIfIndex: vpp2310.NoInterface,
			VrfID:             1,
			Tag:               "DNAT 1",
			Flags:             nat_types.NAT_IS_TWICE_NAT,
		},
		&vpp_nat_ed.Nat44StaticMappingDetails{
			LocalIPAddress:    ipTo4Address("10.10.11.120"),
			ExternalIPAddress: ipTo4Address("10.36.20.31"),
			Protocol:          6,
			LocalPort:         8080,
			ExternalPort:      80,
			ExternalSwIfIndex: vpp2310.NoInterface,
			VrfID:             1,
			Tag:               "DNAT 1",
			Flags:             nat_types.NAT_IS_TWICE_NAT,
		},
		&vpp_nat_ed.Nat44StaticMappingDetails{
			LocalIPAddress:    ipTo4Address("10.10.11.140"),
			ExternalIPAddress: ipTo4Address("10.36.40.10"),
			Protocol:          6,
			LocalPort:         8081,
			ExternalPort:      80,
			ExternalSwIfIndex: vpp2310.NoInterface,
			VrfID:             1,
			Tag:               "DNAT 2",
			Flags:             nat_types.NAT_IS_SELF_TWICE_NAT,
		},
		&vpp_nat_ed.Nat44StaticMappingDetails{
			LocalIPAddress:    ipTo4Address("10.10.11.140"),
			ExternalIPAddress: ipTo4Address("10.36.40.20"),
			Protocol:          6,
			LocalPort:         8081,
			ExternalPort:      80,
			ExternalSwIfIndex: vpp2310.NoInterface,
			VrfID:             1,
			Tag:               "DNAT 2",
			Flags:             nat_types.NAT_IS_SELF_TWICE_NAT,
		},
	)

	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	// LB static mappings
	ctx.MockVpp.MockReply(&vpp_nat_ed.Nat44LbStaticMappingDetails{
		ExternalAddr: ipTo4Address("10.36.20.60"),
		ExternalPort: 53,
		Protocol:     17,
		Flags:        nat_types.NAT_IS_OUT2IN_ONLY,
		Tag:          "DNAT 2",
		LocalNum:     2,
		Locals: []vpp_nat_ed.Nat44LbAddrPort{
			{
				Addr:        ipTo4Address("10.10.11.161"),
				Port:        53,
				Probability: 1,
				VrfID:       0,
			},
			{
				Addr:        ipTo4Address("10.10.11.162"),
				Port:        153,
				Probability: 2,
				VrfID:       0,
			},
		},
	})

	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	// identity mappings
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44IdentityMappingDetails{
			Flags:     nat_types.NAT_IS_ADDR_ONLY,
			Protocol:  17,
			IPAddress: ipTo4Address("10.10.11.200"),
			SwIfIndex: vpp2310.NoInterface,
			VrfID:     1,
			Tag:       "DNAT 3",
		},
		&vpp_nat_ed.Nat44IdentityMappingDetails{
			Flags:     nat_types.NAT_IS_ADDR_ONLY,
			Protocol:  17,
			SwIfIndex: 2,
			VrfID:     1,
			Tag:       "DNAT 3",
		},
		// auto-generated mappings with interface replaced by all assigned IP addresses
		&vpp_nat_ed.Nat44IdentityMappingDetails{
			Flags:     nat_types.NAT_IS_ADDR_ONLY,
			Protocol:  17,
			IPAddress: ipTo4Address("10.36.40.10"),
			SwIfIndex: vpp2310.NoInterface,
			VrfID:     1,
			Tag:       "DNAT 3",
		},
		&vpp_nat_ed.Nat44IdentityMappingDetails{
			Flags:     nat_types.NAT_IS_ADDR_ONLY,
			Protocol:  17,
			IPAddress: ipTo4Address("10.36.40.20"),
			SwIfIndex: vpp2310.NoInterface,
			VrfID:     1,
			Tag:       "DNAT 3",
		},
	)

	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	// interfaces and their IP addresses
	swIfIndexes.Put("if0", &ifaceidx.IfaceMetadata{SwIfIndex: 1, IPAddresses: []string{"10.36.20.30", "10.36.20.31"}})
	swIfIndexes.Put("if1", &ifaceidx.IfaceMetadata{SwIfIndex: 2, IPAddresses: []string{"10.36.40.10"}})
	dhcpIndexes.Put("if1", &ifs.DHCPLease{InterfaceName: "if0", HostIpAddress: "10.36.40.20"})

	dnats, err := natHandler.DNat44Dump()
	Expect(err).To(Succeed())

	Expect(dnats).To(HaveLen(3))

	dnat := dnats[0]
	Expect(dnat.Label).To(Equal("DNAT 1"))
	Expect(dnat.IdMappings).To(HaveLen(0))
	Expect(dnat.StMappings).To(HaveLen(2))
	// 1st mapping
	Expect(dnat.StMappings[0].TwiceNat).To(Equal(nat.DNat44_StaticMapping_ENABLED))
	Expect(dnat.StMappings[0].Protocol).To(Equal(nat.DNat44_TCP))
	Expect(dnat.StMappings[0].ExternalInterface).To(BeEmpty())
	Expect(dnat.StMappings[0].ExternalIp).To(Equal("10.36.20.20"))
	Expect(dnat.StMappings[0].ExternalPort).To(BeEquivalentTo(80))
	Expect(dnat.StMappings[0].LocalIps).To(HaveLen(1))
	Expect(dnat.StMappings[0].LocalIps[0].VrfId).To(BeEquivalentTo(1))
	Expect(dnat.StMappings[0].LocalIps[0].LocalIp).To(Equal("10.10.11.120"))
	Expect(dnat.StMappings[0].LocalIps[0].LocalPort).To(BeEquivalentTo(8080))
	Expect(dnat.StMappings[0].LocalIps[0].Probability).To(BeEquivalentTo(0))
	// 2nd mapping
	Expect(dnat.StMappings[1].TwiceNat).To(Equal(nat.DNat44_StaticMapping_ENABLED))
	Expect(dnat.StMappings[1].Protocol).To(Equal(nat.DNat44_TCP))
	Expect(dnat.StMappings[1].ExternalInterface).To(BeEquivalentTo("if0"))
	Expect(dnat.StMappings[1].ExternalIp).To(BeEquivalentTo("0.0.0.0"))
	Expect(dnat.StMappings[1].ExternalPort).To(BeEquivalentTo(80))
	Expect(dnat.StMappings[1].LocalIps).To(HaveLen(1))
	Expect(dnat.StMappings[1].LocalIps[0].VrfId).To(BeEquivalentTo(1))
	Expect(dnat.StMappings[1].LocalIps[0].LocalIp).To(Equal("10.10.11.120"))
	Expect(dnat.StMappings[1].LocalIps[0].LocalPort).To(BeEquivalentTo(8080))
	Expect(dnat.StMappings[1].LocalIps[0].Probability).To(BeEquivalentTo(0))

	dnat = dnats[1]
	// -> non-LB mapping
	Expect(dnat.Label).To(Equal("DNAT 2"))
	Expect(dnat.IdMappings).To(HaveLen(0))
	Expect(dnat.StMappings).To(HaveLen(2))
	Expect(dnat.StMappings[0].TwiceNat).To(Equal(nat.DNat44_StaticMapping_SELF))
	Expect(dnat.StMappings[0].Protocol).To(Equal(nat.DNat44_TCP))
	Expect(dnat.StMappings[0].ExternalInterface).To(Equal("if1"))
	Expect(dnat.StMappings[0].ExternalIp).To(BeEquivalentTo("0.0.0.0"))
	Expect(dnat.StMappings[0].ExternalPort).To(BeEquivalentTo(80))
	Expect(dnat.StMappings[0].LocalIps).To(HaveLen(1))
	Expect(dnat.StMappings[0].LocalIps[0].VrfId).To(BeEquivalentTo(1))
	Expect(dnat.StMappings[0].LocalIps[0].LocalIp).To(Equal("10.10.11.140"))
	Expect(dnat.StMappings[0].LocalIps[0].LocalPort).To(BeEquivalentTo(8081))
	Expect(dnat.StMappings[0].LocalIps[0].Probability).To(BeEquivalentTo(0))
	// -> LB mapping
	Expect(dnat.StMappings[1].TwiceNat).To(Equal(nat.DNat44_StaticMapping_DISABLED))
	Expect(dnat.StMappings[1].Protocol).To(Equal(nat.DNat44_UDP))
	Expect(dnat.StMappings[1].ExternalInterface).To(BeEmpty())
	Expect(dnat.StMappings[1].ExternalIp).To(Equal("10.36.20.60"))
	Expect(dnat.StMappings[1].ExternalPort).To(BeEquivalentTo(53))
	Expect(dnat.StMappings[1].LocalIps).To(HaveLen(2))
	Expect(dnat.StMappings[1].LocalIps[0].VrfId).To(BeEquivalentTo(0))
	Expect(dnat.StMappings[1].LocalIps[0].LocalIp).To(Equal("10.10.11.161"))
	Expect(dnat.StMappings[1].LocalIps[0].LocalPort).To(BeEquivalentTo(53))
	Expect(dnat.StMappings[1].LocalIps[0].Probability).To(BeEquivalentTo(1))
	Expect(dnat.StMappings[1].LocalIps[1].VrfId).To(BeEquivalentTo(0))
	Expect(dnat.StMappings[1].LocalIps[1].LocalIp).To(Equal("10.10.11.162"))
	Expect(dnat.StMappings[1].LocalIps[1].LocalPort).To(BeEquivalentTo(153))
	Expect(dnat.StMappings[1].LocalIps[1].Probability).To(BeEquivalentTo(2))

	dnat = dnats[2]
	Expect(dnat.Label).To(Equal("DNAT 3"))
	Expect(dnat.StMappings).To(HaveLen(0))
	Expect(dnat.IdMappings).To(HaveLen(2))
	// 1st mapping
	Expect(dnat.IdMappings[0].VrfId).To(BeEquivalentTo(1))
	Expect(dnat.IdMappings[0].Protocol).To(Equal(nat.DNat44_UDP))
	Expect(dnat.IdMappings[0].Port).To(BeEquivalentTo(0))
	Expect(dnat.IdMappings[0].IpAddress).To(Equal("10.10.11.200"))
	Expect(dnat.IdMappings[0].Interface).To(BeEmpty())
	// 2nd mapping
	Expect(dnat.IdMappings[1].VrfId).To(BeEquivalentTo(1))
	Expect(dnat.IdMappings[1].Protocol).To(Equal(nat.DNat44_UDP))
	Expect(dnat.IdMappings[1].Port).To(BeEquivalentTo(0))
	Expect(dnat.IdMappings[1].IpAddress).To(BeEquivalentTo("0.0.0.0"))
	Expect(dnat.IdMappings[1].Interface).To(BeEquivalentTo("if1"))
}

func TestNat44VrfTableDump(t *testing.T) {
	ctx, natHandler, _, _ := natTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vpp_nat_ed.Nat44EdPluginEnableDisableReply{})
	err := natHandler.EnableNAT44Plugin(vppcalls.Nat44InitOpts{EndpointDependent: true})
	Expect(err).ShouldNot(HaveOccurred())

	// vrf tables
	ctx.MockVpp.MockReply(
		&vpp_nat_ed.Nat44EdVrfTablesDetails{
			TableVrfID: 5,
			VrfIds:     []uint32{bits.ReverseBytes32(1), bits.ReverseBytes32(2)},
			NVrfIds:    1,
		},
		&vpp_nat_ed.Nat44EdVrfTablesDetails{
			TableVrfID: 10,
			VrfIds:     []uint32{0},
			NVrfIds:    1,
		})
	ctx.MockVpp.MockReply(&memclnt.ControlPingReply{})

	tables, err := natHandler.Nat44VrfTablesDump()
	Expect(err).To(Succeed())
	Expect(tables).To(HaveLen(2))

	Expect(tables[0].DestVrfIds[0]).To(BeEquivalentTo(1))
	Expect(tables[0].DestVrfIds[1]).To(BeEquivalentTo(2))
	Expect(tables[0].SrcVrfId).To(BeEquivalentTo(5))

	Expect(tables[1].DestVrfIds[0]).To(BeEquivalentTo(0))
	Expect(tables[1].SrcVrfId).To(BeEquivalentTo(10))

}

func natTestSetup(t *testing.T) (*vppmock.TestCtx, vppcalls.NatVppAPI, ifaceidx.IfaceMetadataIndexRW, idxmap.NamedMappingRW) {
	ctx := vppmock.SetupTestCtx(t)
	log := logrus.NewLogger("test-log")
	swIfIndexes := ifaceidx.NewIfaceIndex(logrus.DefaultLogger(), "test-sw_if_indexes")
	dhcpIndexes := idxmap_mem.NewNamedMapping(logrus.DefaultLogger(), "test-dhcp_indexes", nil)
	natHandler := vpp2310.NewNatVppHandler(ctx.MockVPPClient, swIfIndexes, dhcpIndexes, log)
	return ctx, natHandler, swIfIndexes, dhcpIndexes
}

func ipTo4Address(ipStr string) (addr ip_types.IP4Address) {
	netIP := net.ParseIP(ipStr)
	if ip4 := netIP.To4(); ip4 != nil {
		copy(addr[:], ip4)
	}
	return
}
