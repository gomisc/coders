package uuid

import (
	"crypto/sha1"
	"hash"
)

func UUIDv5(base uuid, arg string) []byte {
	return fromHash(sha1.New(), base, arg)
}

func fromHash(h hash.Hash, base uuid, in string) []byte {
	id := make([]byte, 16)
	
	h.Write(base[:])
	h.Write([]byte(in))
	copy(id[:], h.Sum(nil))
	
	return id
}