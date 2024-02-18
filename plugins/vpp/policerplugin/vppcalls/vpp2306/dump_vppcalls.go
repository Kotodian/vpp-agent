package vpp2310

import (
	vpp_policer "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/policer"
	policer "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"
)

// DumpPolicers implements policer handler.
func (h *PolicerVppHandler) DumpPolicers() (policerList []*policer.PolicerConfig, err error) {
	// index of ^uint32(0) dumps all peers
	req := &vpp_policer.PolicerDumpV2{PolicerIndex: ^uint32(0)}
	requestCtx := h.callsChannel.SendMultiRequest(req)

	var vppPolicerList []*vpp_policer.PolicerDetails
	for {
		vppPolicerDetails := &vpp_policer.PolicerDetails{}
		stop, err := requestCtx.ReceiveReply(vppPolicerDetails)
		if stop {
			break
		}
		if err != nil {
			return nil, err
		}
		vppPolicerList = append(vppPolicerList, vppPolicerDetails)
	}

	for _, vppPolicerDetails := range vppPolicerList {
		policerDetails := &policer.PolicerConfig{
			Name:       vppPolicerDetails.Name,
			Cir:        vppPolicerDetails.Cir,
			Eir:        vppPolicerDetails.Eir,
			Cb:         vppPolicerDetails.Cb,
			Eb:         vppPolicerDetails.Eb,
			RateType:   policer.Sse2QosRateType(vppPolicerDetails.RateType),
			RoundType:  policer.Sse2QosRoundType(vppPolicerDetails.RoundType),
			Type:       policer.Sse2QosPolicerType(vppPolicerDetails.Type),
			ColorAware: vppPolicerDetails.ColorAware,
			ConformAction: &policer.Sse2QosAction{
				Type: policer.Sse2QosActionType(vppPolicerDetails.ConformAction.Type),
				Dscp: uint32(vppPolicerDetails.ConformAction.Dscp),
			},
			ExceedAction: &policer.Sse2QosAction{
				Type: policer.Sse2QosActionType(vppPolicerDetails.ExceedAction.Type),
				Dscp: uint32(vppPolicerDetails.ExceedAction.Dscp),
			},
			ViolateAction: &policer.Sse2QosAction{
				Type: policer.Sse2QosActionType(vppPolicerDetails.ViolateAction.Type),
				Dscp: uint32(vppPolicerDetails.ViolateAction.Dscp),
			},
		}

		policerList = append(policerList, policerDetails)
	}

	return
}
