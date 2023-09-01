package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeviceKeyPrefix is the prefix to retrieve all Device
	DeviceKeyPrefix = "Device/value/"
)

// DeviceKey returns the store key to retrieve a Device from the index fields
func DeviceKey(
	deviceName string,
) []byte {
	var key []byte

	deviceNameBytes := []byte(deviceName)
	key = append(key, deviceNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
