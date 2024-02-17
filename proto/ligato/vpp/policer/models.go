package vpp_policer

import (
	"strings"

	"go.ligato.io/vpp-agent/v3/pkg/models"
)

// ModuleName is the module name used for models.
const ModuleName = "vpp.policer"

var (
	ModelPolicerConfig models.KnownModel
)

func init() {
	// models.Register requires protoreflect capabilities, so we initialize them first
	file_ligato_vpp_policer_policer_proto_init()

	ModelPolicerConfig = models.Register(&PolicerConfig{}, models.Spec{
		Module:  ModuleName,
		Type:    "config",
		Version: "v2",
	}, models.WithNameTemplate("{{.Name}}"))
}

func PolicerConfigKey(name string) string {
	return models.Key(&PolicerConfig{
		Name: name,
	})
}

/* Policer interface (derived) */

const (
	// interfacePolicerKeyPrefix is a common prefix for (derived) keys each representing
	// policer configuration for a single interface.
	interfacePolicerKeyPrefix = "vpp/policer/"

	// interfacePolicerKeyTemplate is a template for (derived) key representing
	// Policer configuration for a single interface.
	interfacePolicerKeyTemplate = interfacePolicerKeyPrefix + "{name}/interface/{iface}/feature/{feature}"

	// Policer interface features
	inputFeature  = "input"
	outputFeature = "output"
)

const (
	// InvalidKeyPart is used in key for parts which are invalid
	InvalidKeyPart = "<invalid>"
)

// DerivedPolicerInterfaceKey returns (derived) key representing Policer configuration
// for a given interface.
func DerivedPolicerInterfaceKey(policerName, iface string, isOutput bool) string {
	if iface == "" {
		iface = InvalidKeyPart
	}
	key := strings.Replace(interfacePolicerKeyTemplate, "{name}", policerName, 1)
	key = strings.Replace(key, "{iface}", iface, 1)
	feature := inputFeature
	if isOutput {
		feature = outputFeature
	}
	key = strings.Replace(key, "{feature}", feature, 1)
	return key
}

// ParseDerivedInterfacePolicerKey parses interface name and the assigned Policer feature
// from Interface-Policer key.
func ParseDerivedPolicerInterfaceKey(key string) (policerName, iface string, isOutput bool, isInterfacePolicerKey bool) {
	trim := strings.TrimPrefix(key, interfacePolicerKeyPrefix)
	if trim != key && trim != "" {
		fibComps := strings.Split(trim, "/")
		if len(fibComps) >= 5 && fibComps[1] == "interface" && fibComps[len(fibComps)-2] == "feature" {
			if fibComps[len(fibComps)-1] == outputFeature {
				isOutput = true
			} else if fibComps[len(fibComps)-1] != inputFeature {
				return
			}
			policerName = fibComps[0]
			fibComps = fibComps[2:]
			iface = strings.Join(fibComps[:len(fibComps)-2], "/")
			isInterfacePolicerKey = true
			return
		}
	}
	return "", "", false, false
}
