package vpp2310

import (
	"fmt"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/interface_types"
	vpp_policer "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/policer"
	vpp_policer_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/policer_types"
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

func (h *PolicerVppHandler) UpdatePolicer(policerIndex uint32, cfg *policer.PolicerConfig) error {
	policerVpp := vpp_policer_types.PolicerConfig{
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
		policerVpp.ConformAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policerVpp.ConformAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ConformAction.Type),
			Dscp: uint8(cfg.ConformAction.Dscp),
		}
	}

	if cfg.ExceedAction == nil {
		policerVpp.ExceedAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policerVpp.ExceedAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ExceedAction.Type),
			Dscp: uint8(cfg.ExceedAction.Dscp),
		}
	}

	if cfg.ViolateAction == nil {
		policerVpp.ViolateAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.SSE2_QOS_ACTION_API_DROP,
			Dscp: 0,
		}
	} else {
		policerVpp.ViolateAction = vpp_policer_types.Sse2QosAction{
			Type: vpp_policer_types.Sse2QosActionType(cfg.ViolateAction.Type),
			Dscp: uint8(cfg.ViolateAction.Dscp),
		}
	}

	request := &vpp_policer.PolicerUpdate{
		PolicerIndex: policerIndex,
		Infos:        policerVpp,
	}
	// prepare reply
	reply := &vpp_policer.PolicerUpdateReply{}
	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}

	return nil
}

func (h *PolicerVppHandler) DelPolicer(policerIndex uint32) error {
	// prepare request
	request := &vpp_policer.PolicerDel{
		PolicerIndex: policerIndex,
	}
	// prepare reply
	reply := &vpp_policer.PolicerDelReply{}

	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil
}

func (h *PolicerVppHandler) ResetPolicer(policerIndex uint32) error {
	// prepare request
	request := &vpp_policer.PolicerReset{
		PolicerIndex: policerIndex,
	}
	// prepare reply
	reply := &vpp_policer.PolicerResetReply{}

	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil
}

func (h *PolicerVppHandler) PolicerInput(policerIndex uint32, iface *policer.PolicerConfig_Interface, apply bool) error {
	im, exists := h.ifIdx.LookupByName(iface.Name)
	if !exists {
		return fmt.Errorf("interface: %s doesn't exist", iface.Name)
	}
	request := &vpp_policer.PolicerInputV2{
		SwIfIndex:    interface_types.InterfaceIndex(im.SwIfIndex),
		PolicerIndex: policerIndex,
		Apply:        apply,
	}
	// prepare reply, do not use PolicerInputV2Reply, the reason i think is vpp's api bug
	reply := &vpp_policer.PolicerInputReply{}
	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil
}

func (h *PolicerVppHandler) PolicerOutput(policerIndex uint32, iface *policer.PolicerConfig_Interface, apply bool) error {
	im, exists := h.ifIdx.LookupByName(iface.Name)
	if !exists {
		return fmt.Errorf("interface: %s doesn't exist", iface.Name)
	}
	request := &vpp_policer.PolicerOutputV2{
		SwIfIndex:    interface_types.InterfaceIndex(im.SwIfIndex),
		PolicerIndex: policerIndex,
		Apply:        apply,
	}
	// prepare reply, do not use PolicerOutputV2Reply, the reason i think is vpp's api bug
	reply := &vpp_policer.PolicerOutputReply{}
	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil
}

func (h *PolicerVppHandler) PolicerBind(policerIndex uint32, worker *policer.PolicerConfig_Worker, enable bool) error {
	request := &vpp_policer.PolicerBindV2{
		PolicerIndex: policerIndex,
		WorkerIndex:  worker.Index,
		BindEnable:   enable,
	}
	// prepare reply
	reply := &vpp_policer.PolicerBindV2Reply{}
	// send request and obtain reply
	if err := h.callsChannel.SendRequest(request).ReceiveReply(reply); err != nil {
		return err
	}
	return nil
}
