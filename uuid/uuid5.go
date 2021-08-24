package uuid

import (
	"crypto/sha1"
	"hash"
	"math/big"
)

// UUIDv5 - возвращает новый UUID5 в байтах
func UUIDv5(base uuid, arg string) []byte {
	return fromHash(sha1.New(), base, arg)
}

// IntUUIDv5 - возвращает новый UUID5 в big.Int
func IntUUIDv5(base uuid, arg string) *big.Int {
	id := UUIDv5(base, arg)
	z := big.NewInt(0)
	
	return z.SetBytes(id)
}

// StrUUIDv5 - возвращает новый UUID5 строкой
func StrUUIDv5(base uuid, arg string) string {
	id := UUIDv5(base, arg)
	
	return idBytesToStr(id)
}

func fromHash(h hash.Hash, base uuid, in string) []byte {
	id := make([]byte, 16)
	
	h.Write(base[:])
	h.Write([]byte(in))
	copy(id[:], h.Sum(nil))
	
	return id
}