// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"google.golang.org/protobuf/proto"
	. "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/policerplugin/policeridx"
	"go.ligato.io/vpp-agent/v3/proto/ligato/vpp/policer"
)

////////// type-safe key-value pair with metadata //////////

type PolicerKVWithMetadata struct {
	Key      string
	Value    *vpp_policer.PolicerConfig
	Metadata *policeridx.PolicerMetadata
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type PolicerDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *vpp_policer.PolicerConfig) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *vpp_policer.PolicerConfig) error
	Create               func(key string, value *vpp_policer.PolicerConfig) (metadata *policeridx.PolicerMetadata, err error)
	Delete               func(key string, value *vpp_policer.PolicerConfig, metadata *policeridx.PolicerMetadata) error
	Update               func(key string, oldValue, newValue *vpp_policer.PolicerConfig, oldMetadata *policeridx.PolicerMetadata) (newMetadata *policeridx.PolicerMetadata, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *vpp_policer.PolicerConfig, metadata *policeridx.PolicerMetadata) bool
	Retrieve             func(correlate []PolicerKVWithMetadata) ([]PolicerKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *vpp_policer.PolicerConfig) []KeyValuePair
	Dependencies         func(key string, value *vpp_policer.PolicerConfig) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type PolicerDescriptorAdapter struct {
	descriptor *PolicerDescriptor
}

func NewPolicerDescriptor(typedDescriptor *PolicerDescriptor) *KVDescriptor {
	adapter := &PolicerDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *PolicerDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castPolicerValue(key, oldValue)
	typedNewValue, err2 := castPolicerValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *PolicerDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castPolicerValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *PolicerDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castPolicerValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *PolicerDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castPolicerValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castPolicerValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castPolicerMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *PolicerDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castPolicerValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castPolicerMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *PolicerDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castPolicerValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castPolicerValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castPolicerMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *PolicerDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []PolicerKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castPolicerValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castPolicerMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			PolicerKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *PolicerDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castPolicerValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *PolicerDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castPolicerValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castPolicerValue(key string, value proto.Message) (*vpp_policer.PolicerConfig, error) {
	typedValue, ok := value.(*vpp_policer.PolicerConfig)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castPolicerMetadata(key string, metadata Metadata) (*policeridx.PolicerMetadata, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(*policeridx.PolicerMetadata)
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}