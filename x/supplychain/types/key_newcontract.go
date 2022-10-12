package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NewcontractKeyPrefix is the prefix to retrieve all Newcontract
	NewcontractKeyPrefix = "Newcontract/value/"
)

// NewcontractKey returns the store key to retrieve a Newcontract from the index fields
func NewcontractKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
