package uuid

import (
	"bytes"
	"encoding/hex"
	"strings"
)

const Size = 16

type uuid [Size]byte

var Zero = uuid{}

// UUID возвращает UUIDv5 если переданы аргументы, либо UUIDv4
// в байтах
func UUID(from ...string) []byte {
	if len(from) != 0 {
		return UUIDv5(Zero, strings.Join(from, ""))
	}

	return UUIDv4()
}

// StrUUID возвращает UUIDv5 если переданы аргументы, либо UUIDv4
// стройкой
func StrUUID(from ...string) string {
	return idBytesToStr(UUID(from...))
}

// FromString - возвращает UUID в байтах по строковому представлению
func FromString(s string) ([]byte, error) {
	id := make([]byte, Size)

	hexStr := strings.ReplaceAll(s, "-", "")
	if _, err := hex.Decode(id, []byte(hexStr)); err != nil {
		return nil, err
	}

	return id, nil
}

// FromBytes - возвращает UUID строкой из представления в байтах
func FromBytes(u []byte) string {
	return idBytesToStr(u)
}

// форматирует UUID4 в байтах в строку
func idBytesToStr(id []byte) string {
	const delim = "-"

	var buf bytes.Buffer

	buf.WriteString(hex.EncodeToString(id[0:4]))
	buf.WriteString(delim)

	buf.WriteString(hex.EncodeToString(id[4:6]))
	buf.WriteString(delim)

	buf.WriteString(hex.EncodeToString(id[6:8]))
	buf.WriteString(delim)

	buf.WriteString(hex.EncodeToString(id[8:10]))
	buf.WriteString(delim)

	buf.WriteString(hex.EncodeToString(id[10:]))
	buf.WriteString(delim)

	return buf.String()
}
