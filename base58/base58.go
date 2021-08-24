package base58

// Alphabet - base58 alphabet
type Alphabet interface {
	EncodeSymbols() [58]byte
	DecodeSymbols() [128]int8
}

// Encoder - base58 encoder
type Encoder interface {
	Encode(in []byte) string
	Decode(str string) ([]byte, error)
}

// Encode - encode base58 with default encoder
func Encode(in []byte) string {
	return defaultEncoder.Encode(in)
}

// Decode - decode base58 with default encoder
func Decode(in string) ([]byte, error) {
	return defaultEncoder.Decode(in)
}