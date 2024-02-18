package policerplugin

import (
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/policeridx"
)

// API defines methods exposed by VPP-PolicerPlugin.
type API interface {
	// GetPolicerIndex gives read-only access to map with metadata of all configured
	// VPP policer lists.
	GetPolicerIndex() policeridx.PolicerMetadataIndex
}
