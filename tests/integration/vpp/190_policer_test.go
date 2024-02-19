package vpp

import (
	"testing"

	. "github.com/onsi/gomega"
	"go.ligato.io/cn-infra/v2/logging/logrus"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/ifaceidx"
	ifplugin_vppcalls "go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/vppcalls"
	policerplugin_vppcalls "go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/vppcalls"
	policer "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"

	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/aclplugin"
	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin"
)

var (
	policer1 = &policer.PolicerConfig{
		Name:      "policer1",
		Cir:       100,
		Eir:       0,
		Cb:        100000,
		Eb:        0,
		RateType:  policer.Sse2QosRateType_RATE_KBPS,
		RoundType: policer.Sse2QosRoundType_ROUND_TO_CLOSEST,
		Type:      policer.Sse2QosPolicerType_POLICER_1R2C,
		ConformAction: &policer.Sse2QosAction{
			Type: policer.Sse2QosActionType_ACTION_TRANSMIT,
		},
		ExceedAction: &policer.Sse2QosAction{
			Type: policer.Sse2QosActionType_ACTION_DROP,
		},
		ViolateAction: &policer.Sse2QosAction{
			Type: policer.Sse2QosActionType_ACTION_DROP,
		},
	}
)

func TestPolicerCURD(t *testing.T) {
	ctx := setupVPP(t)
	defer ctx.teardownVPP()

	ih := ifplugin_vppcalls.CompatibleInterfaceVppHandler(ctx.vppClient, logrus.NewLogger("test"))
	Expect(ih).To(Not(BeNil()), "Handler should be created.")

	const ifName = "loop1"
	ifIdx, errI := ih.AddLoopbackInterface(ifName)
	Expect(errI).To(BeNil())
	t.Logf("Prerequsite: interface %v created - its index %v", ifName, ifIdx)

	ifIndexes := ifaceidx.NewIfaceIndex(logrus.NewLogger("test-iface1"), "test-iface1")
	ifIndexes.Put(ifName, &ifaceidx.IfaceMetadata{
		SwIfIndex: ifIdx,
	})

	h := policerplugin_vppcalls.CompatiblePolicerVppHandler(ctx.vppClient, ifIndexes, logrus.NewLogger("policer-vppcalls"))
	Expect(h).To(Not(BeNil()), "Handler should be created.")

	policers, err := h.DumpPolicers()
	Expect(err).To(BeNil())
	Expect(policers).Should(BeEmpty())
	t.Log("no policers dumped")

	// create POLICER
	policerIdx, err := h.AddPolicer(policer1)
	Expect(err).To(BeNil())
	Expect(policerIdx).To(BeEquivalentTo(0))
	t.Logf("policer \"%v\" added - its index %d", policer1.Name, policerIdx)

	// update POLICER interface
	policer1.Interfaces = append(policer1.Interfaces, &policer.PolicerConfig_Interface{
		Name:     ifName,
		IsOutput: true,
	})
	Expect(h.PolicerInput(policerIdx, policer1.Interfaces[0], true)).To(Succeed())

	policers, err = h.DumpPolicers()
	Expect(err).To(BeNil())
	Expect(policers).Should(HaveLen(1))

	// update POLICER
	policer1.Cb = 1000
	Expect(h.UpdatePolicer(policerIdx, policer1)).To(Succeed())

	// update POLICER interface
	Expect(h.PolicerInput(policerIdx, policer1.Interfaces[0], false)).To(Succeed())
	Expect(h.PolicerOutput(policerIdx, policer1.Interfaces[0], true)).To(Succeed())
	Expect(h.PolicerOutput(policerIdx, policer1.Interfaces[0], false)).To(Succeed())
	
	// delete POLICER
	Expect(h.DelPolicer(policerIdx)).To(Succeed())
}
