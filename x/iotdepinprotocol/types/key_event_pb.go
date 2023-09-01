package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EventPbKeyPrefix is the prefix to retrieve all EventPb
	EventPbKeyPrefix = "EventPb/value/"
)

// EventPbKey returns the store key to retrieve a EventPb from the index fields
func EventPbKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
