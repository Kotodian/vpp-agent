package vpp_policer

import "go.ligato.io/vpp-agent/v3/pkg/models"

// ModuleName is the module name used for models.
const ModuleName = "vpp"

var (
	ModelPolicerConfig models.KnownModel
)

func init() {
	// models.Register requires protoreflect capabilities, so we initialize them first
	file_ligato_vpp_policer_policer_proto_init()

	ModelPolicerConfig = models.Register(&PolicerConfig{}, models.Spec{
		Module:  ModuleName,
		Type:    "policerconfig",
		Version: "v2",
	}, models.WithNameTemplate("{{.Name}}"))
}

func PolicerConfigKey(name string) string {
	return models.Key(&PolicerConfig{
		Name: name,
	})
}
