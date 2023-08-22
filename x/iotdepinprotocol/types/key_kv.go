package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// KvKeyPrefix is the prefix to retrieve all Kv
	KvKeyPrefix = "Kv/value/"
)

// KvKey returns the store key to retrieve a Kv from the index fields
func KvKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
