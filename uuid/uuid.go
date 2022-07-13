package uuid

import (
	"bytes"
	"encoding/hex"
	"strings"

	"git.corout.in/golibs/errors"
)

const (
	// Size - размер UUID
	Size = 16

	v4 = byte(4)
	v5 = byte(5)
)

type uuid [Size]byte

var Zero = uuid{} // nolint: gochecknoglobals

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

func setVersion(id []byte, v byte) {
	// вариант https://datatracker.ietf.org/doc/html/rfc4122#section-4.1.1
	id[6] = (id[6] & 0x0f) | (v << 4)
	// версия https://datatracker.ietf.org/doc/html/rfc4122#section-4.1.3
	id[8] = id[8]&(0xff>>2) | (0x02 << 6)
}

// форматирует UUID4 в байтах в строку
func idBytesToStr(id []byte) string {
	const delim = "-"

	var buf bytes.Buffer

	err := writeStrings(&buf,
		hex.EncodeToString(id[0:4]),
		delim,
		hex.EncodeToString(id[4:6]),
		delim,
		hex.EncodeToString(id[6:8]),
		delim,
		hex.EncodeToString(id[8:10]),
		delim,
		hex.EncodeToString(id[10:]),
	)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func writeStrings(buf *bytes.Buffer, strs ...string) error {
	for _, str := range strs {
		if _, err := buf.WriteString(str); err != nil {
			return errors.Wrap(err, "write into buffer")
		}
	}

	return nil
}
