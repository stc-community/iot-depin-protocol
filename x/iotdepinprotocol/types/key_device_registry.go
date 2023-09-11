package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeviceRegistryKeyPrefix is the prefix to retrieve all DeviceRegistry
	DeviceRegistryKeyPrefix = "DeviceRegistry/value/"
)

// DeviceRegistryKey returns the store key to retrieve a DeviceRegistry from the index fields
func DeviceRegistryKey(
	mid string,
) []byte {
	var key []byte

	midBytes := []byte(mid)
	key = append(key, midBytes...)
	key = append(key, []byte("/")...)

	return key
}
