package uuid

import (
	"crypto/sha1" // nolint: gosec
	"hash"
	"math/big"
)

// UUIDv5 - возвращает новый UUID5 в байтах
// nolint: gosec
func UUIDv5(ns uuid, arg string) []byte {
	return fromHash(sha1.New(), ns, arg)
}

// IntUUIDv5 - возвращает новый UUID5 в big.Int
func IntUUIDv5(ns uuid, arg string) *big.Int {
	id := UUIDv5(ns, arg)
	z := big.NewInt(0)

	return z.SetBytes(id)
}

// StrUUIDv5 - возвращает новый UUID5 строкой
func StrUUIDv5(ns uuid, arg string) string {
	id := UUIDv5(ns, arg)

	return idBytesToStr(id)
}

func fromHash(h hash.Hash, ns uuid, in string) []byte {
	id := make([]byte, 16)

	h.Write(ns[:])
	h.Write([]byte(in))

	copy(id, h.Sum(nil))

	setVersion(id, v5)

	return id
}
