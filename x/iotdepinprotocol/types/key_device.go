package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeviceKeyPrefix is the prefix to retrieve all Device
	DeviceKeyPrefix = "Device/value/"
)

// DeviceKey returns the store key to retrieve a Device from the index fields
func DeviceKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
