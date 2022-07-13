package uuid

import (
	"crypto/rand"
	"math/big"
)

// UUIDv4 - возвращает новый UUID4 в байтах
func UUIDv4() []byte {
	id := make([]byte, Size)

	if _, err := rand.Read(id); err != nil {
		panic(err)
	}

	setVersion(id, v4)

	return id
}

// IntUUIDv4 - возвращает новый UUID4 в big.Int
func IntUUIDv4() *big.Int {
	id := UUIDv4()
	z := big.NewInt(0)

	return z.SetBytes(id)
}

// StrUUIDv4 - возвращает новый UUID4 строкой
func StrUUIDv4() string {
	id := UUIDv4()

	return idBytesToStr(id)
}
