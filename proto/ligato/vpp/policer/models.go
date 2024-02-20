package vpp_policer

import (
	"strconv"
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

const (
	// interfacePolicerKeyPrefix is a common prefix for (derived) keys each representing
	// policer configuration for a single interface.
	policerKeyPrefix = "vpp/policer/"

	// interfacePolicerKeyTemplate is a template for (derived) key representing
	// Policer configuration for a single interface.
	interfacePolicerKeyTemplate = policerKeyPrefix + "{name}/interface/{iface}/feature/{feature}"

	// Policer interface features
	inputFeature  = "input"
	outputFeature = "output"

	// workerPolicerKeyTemplate is a template for (derived) key representing
	// Policer configuration for a worker thread.
	workerPolicerKeyTemplate = policerKeyPrefix + "{name}/worker/{worker}"
)

const (
	// InvalidKeyPart is used in key for parts which are invalid
	InvalidKeyPart = "<invalid>"
)

/* Policer interface (derived) */

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
	trim := strings.TrimPrefix(key, policerKeyPrefix)
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

/* Policer worker (derived) */

// DerivedPolicerWorkerKey returns (derived) key representing Policer configuration
// for a given worker thread.
func DerivedPolicerWorkerKey(policerName string, workerIndex uint32) string {
	key := strings.Replace(workerPolicerKeyTemplate, "{name}", policerName, 1)
	if workerIndex == ^uint32(0) {
		key = strings.Replace(key, "{worker}", InvalidKeyPart, 1)
	} else {
		key = strings.Replace(key, "{worker}", strconv.Itoa(int(workerIndex)), 1)
	}
	return key
}

// ParseDerivedPolicerWorkerKey parses policer name and the assigned Policer worker
func ParseDerivedPolicerWorkerKey(key string) (policerName string, workerIndex uint32, isPolicerWorkerKey bool) {
	workerIndex = ^uint32(0)
	trim := strings.TrimPrefix(key, policerKeyPrefix)
	if trim != key && trim != "" {
		fibComps := strings.Split(trim, "/")
		if len(fibComps) >= 3 && fibComps[1] == "worker" {
			policerName = fibComps[0]
			if fibComps[2] == InvalidKeyPart {
				return
			}
			_workerIndex, err := strconv.ParseUint(fibComps[2], 10, 32)
			if err != nil {
				return
			}
			workerIndex = uint32(_workerIndex)
			isPolicerWorkerKey = true
			return
		}
	}
	return
}
