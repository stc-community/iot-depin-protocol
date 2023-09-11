package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DeviceList:         []Device{},
		KvList:             []Kv{},
		EventPbList:        []EventPb{},
		DeviceRegistryList: []DeviceRegistry{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in device
	deviceIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeviceList {
		index := string(DeviceKey(elem.DeviceName))
		if _, ok := deviceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for device")
		}
		deviceIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in kv
	kvIndexMap := make(map[string]struct{})

	for _, elem := range gs.KvList {
		index := string(KvKey(elem.Index))
		if _, ok := kvIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for kv")
		}
		kvIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in eventPb
	eventPbIndexMap := make(map[string]struct{})

	for _, elem := range gs.EventPbList {
		index := string(EventPbKey(elem.Index))
		if _, ok := eventPbIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for eventPb")
		}
		eventPbIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in deviceRegistry
	deviceRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeviceRegistryList {
		index := string(DeviceRegistryKey(elem.Mid))
		if _, ok := deviceRegistryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for deviceRegistry")
		}
		deviceRegistryIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
