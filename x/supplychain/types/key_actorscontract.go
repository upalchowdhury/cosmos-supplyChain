package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActorscontractKeyPrefix is the prefix to retrieve all Actorscontract
	ActorscontractKeyPrefix = "Actorscontract/value/"
)

// ActorscontractKey returns the store key to retrieve a Actorscontract from the index fields
func ActorscontractKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
