package vpp2310

import (
	"fmt"

	vpp_policer "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/policer"
	vpp_policer_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/policer_types"
	policer "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"
)

func (h *PolicerVppHandler) AddPolicer(cfg *policer.PolicerConfig) (uint32, error) {
	invalidIdx := ^uint32(0)

	policer_vpp := vpp_policer_types.PolicerConfig{
		Cir:        cfg.Cir,
		Eir:        cfg.Eir,
		Cb:         cfg.Cb,
		Eb:         cfg.Eb,
		RateType:   vpp_policer_types.Sse2QosRateType(cfg.RateType),
		RoundType:  vpp_policer_types.Sse2QosRoundType(cfg.RoundType),
		Type:       vpp_policer_types.Sse2QosPolicerType(cfg.Type),
		ColorAware: cfg.ColorAware,
	}

	if cfg.ConformAction == nil {
		policer_vpp.ConformAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policer_vpp.ConformAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ConformAction.Type),
			Dscp: uint8(cfg.ConformAction.Dscp),
		}
	}

	if cfg.ExceedAction == nil {
		policer_vpp.ExceedAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policer_vpp.ExceedAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ExceedAction.Type),
			Dscp: uint8(cfg.ExceedAction.Dscp),
		}
	}

	if cfg.ViolateAction == nil {
		policer_vpp.ViolateAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policer_vpp.ViolateAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ViolateAction.Type),
			Dscp: uint8(cfg.ViolateAction.Dscp),
		}
	}

	if len([]byte(cfg.Name)) > 64 {
		return invalidIdx, fmt.Errorf("name is invalid")
	}

	request := &vpp_policer.PolicerAdd{
		Name:  cfg.Name,
		Infos: policer_vpp,
	}
	// prepare reply
	reply := &vpp_policer.PolicerAddReply{}
	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return invalidIdx, err
	}

	return reply.PolicerIndex, nil
}

func (h *PolicerVppHandler) UpdatePolicer(policer_index uint32, cfg *policer.PolicerConfig) error {
	policer_vpp := vpp_policer_types.PolicerConfig{
		Cir:        cfg.Cir,
		Eir:        cfg.Eir,
		Cb:         cfg.Cb,
		Eb:         cfg.Eb,
		RateType:   vpp_policer_types.Sse2QosRateType(cfg.RateType),
		RoundType:  vpp_policer_types.Sse2QosRoundType(cfg.RoundType),
		Type:       vpp_policer_types.Sse2QosPolicerType(cfg.Type),
		ColorAware: cfg.ColorAware,
	}

	if cfg.ConformAction == nil {
		policer_vpp.ConformAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policer_vpp.ConformAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ConformAction.Type),
			Dscp: uint8(cfg.ConformAction.Dscp),
		}
	}

	if cfg.ExceedAction == nil {
		policer_vpp.ExceedAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policer_vpp.ExceedAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ExceedAction.Type),
			Dscp: uint8(cfg.ExceedAction.Dscp),
		}
	}

	if cfg.ViolateAction == nil {
		policer_vpp.ViolateAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policer_vpp.ViolateAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ViolateAction.Type),
			Dscp: uint8(cfg.ViolateAction.Dscp),
		}
	}

	request := &vpp_policer.PolicerUpdate{
		PolicerIndex: policer_index,
		Infos:        policer_vpp,
	}
	// prepare reply
	reply := &vpp_policer.PolicerUpdateReply{}
	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}

	return nil
}

func (h *PolicerVppHandler) DelPolicer(policer_idx uint32) error {
	// prepare request
	request := &vpp_policer.PolicerDel{
		PolicerIndex: policer_idx,
	}
	// prepare reply
	reply := &vpp_policer.PolicerDelReply{}

	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil
}
