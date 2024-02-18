package policeridx

import (
	"time"

	"go.ligato.io/cn-infra/v2/idxmap"
	"go.ligato.io/cn-infra/v2/logging"
	"go.ligato.io/vpp-agent/v3/pkg/idxvpp"
)

// PolicerMetadataIndex provides read-only access to mapping between Policer indices (used internally in VPP)
// and Policer names.
type PolicerMetadataIndex interface {
	// LookupByName looks up previously stored item identified by name in mapping.
	LookupByName(name string) (metadata *PolicerMetadata, exists bool)

	// LookupByIdx looks up previously stored item identified by index in mapping.
	LookupByIndex(idx uint32) (name string, metadata *PolicerMetadata, exists bool)

	// WatchPolicers
	WatchPolicers(subscriber string, channel chan<- PolicerMetadataDto)
}

type PolicerMetadataIndexRW interface {
	PolicerMetadataIndex
	idxmap.NamedMappingRW
}

// PolicerMetadata represents metadata for policer
type PolicerMetadata struct {
	Index uint32
}

func (m *PolicerMetadata) GetIndex() uint32 {
	return m.Index
}

// PolicerMetadataDto represents an item sent through watch channel in policerIndex.
type PolicerMetadataDto struct {
	idxmap.NamedMappingEvent
	Metadata *PolicerMetadata
}

type policerMetadataIndex struct {
	idxmap.NamedMappingRW

	log         logging.Logger
	nameToIndex idxvpp.NameToIndex
}

func NewPolicerIndex(logger logging.Logger, title string) PolicerMetadataIndexRW {
	mapping := idxvpp.NewNameToIndex(logger, title, indexMetadata)
	return &policerMetadataIndex{
		NamedMappingRW: mapping,
		log:            logger,
		nameToIndex:    mapping,
	}
}

// LookupByName looks up previously stored item identified by index in mapping.
func (policerIdx *policerMetadataIndex) LookupByName(name string) (metadata *PolicerMetadata, exists bool) {
	meta, found := policerIdx.GetValue(name)
	if found {
		if typedMeta, ok := meta.(*PolicerMetadata); ok {
			return typedMeta, found
		}
	}
	return nil, false
}

// LookupByIndex looks up previously stored item identified by name in mapping.
func (policerIdx *policerMetadataIndex) LookupByIndex(idx uint32) (name string, metadata *PolicerMetadata, exists bool) {
	var item idxvpp.WithIndex
	name, item, exists = policerIdx.nameToIndex.LookupByIndex(idx)
	if exists {
		var isIfaceMeta bool
		metadata, isIfaceMeta = item.(*PolicerMetadata)
		if !isIfaceMeta {
			exists = false
		}
	}
	return
}

func (policerIdx *policerMetadataIndex) WatchPolicers(subscriber string, channel chan<- PolicerMetadataDto) {
	watcher := func(dto idxmap.NamedMappingGenericEvent) {
		typedMeta, ok := dto.Value.(*PolicerMetadata)
		if !ok {
			return
		}
		msg := PolicerMetadataDto{
			NamedMappingEvent: dto.NamedMappingEvent,
			Metadata:          typedMeta,
		}
		select {
		case channel <- msg:
		case <-time.After(idxmap.DefaultNotifTimeout):
			policerIdx.log.Warn("Unable to deliver notification")
		}
	}
	if err := policerIdx.Watch(subscriber, watcher); err != nil {
		policerIdx.log.Error(err)
	}
}

// indexMetadata is an index function used for Policer metadata.
func indexMetadata(metaData interface{}) map[string][]string {
	indexes := make(map[string][]string)

	ifMeta, ok := metaData.(*PolicerMetadata)
	if !ok || ifMeta == nil {
		return indexes
	}

	return indexes
}
