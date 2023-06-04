package base58

import (
	"fmt"

	"gopkg.in/gomisc/errors.v1"
)

const (
	ErrZeroInputLength = errors.Const("input data has zero length")
	ErrHighBit         = errors.Const("high-bit set on invalid digit")
	ErrOutputNumber    = errors.Const("output number too big")
)

var defaultEncoder = New(BTCAlphabet()) // nolint: gochecknoglobals

type encoder struct {
	alphabet Alphabet
}

func New(alphabet Alphabet) Encoder {
	return &encoder{
		alphabet: alphabet,
	}
}

func (e *encoder) Encode(in []byte) string {
	encode := e.alphabet.EncodeSymbols()

	var (
		i, j, zcount, high int
		carry              uint32
		binsz              = len(in)
		zero               = encode[0]
	)

	// расчет количества лидирующих нулей данных для кодирования
	for zcount < binsz && in[zcount] == 0 {
		zcount++
	}

	// расчет размера слайса для хранения преобразованных данных
	size := (binsz-zcount)*138/100 + 1
	buf := make([]byte, size*2+zcount)
	tmp := buf[size+zcount:]
	high = size - 1

	// преобразование данных
	for i = zcount; i < binsz; i++ {
		j = size - 1

		// кодировка последовательности от начала до конца
		for carry = uint32(in[i]); j > high || carry != 0; j-- {
			carry += 256 * uint32(tmp[j])
			tmp[j] = byte(carry % 58)
			carry /= 58
		}

		high = j
	}

	for j = 0; j < size && tmp[j] == 0; j++ {
	}

	// кодирование лидирующих нолей
	b58 := buf[:size-j+zcount]

	if zcount != 0 {
		for i = 0; i < zcount; i++ {
			b58[i] = zero
		}
	}

	for i = zcount; j < size; i++ {
		b58[i] = encode[tmp[j]]
		j++
	}

	return string(b58)
}

func (e *encoder) Decode(in string) ([]byte, error) {
	if in == "" {
		return nil, ErrZeroInputLength
	}

	encode, decode := e.alphabet.EncodeSymbols(), e.alphabet.DecodeSymbols()

	var (
		t, c   uint64
		zmask  uint32
		zcount int

		b58u  = []rune(in)
		b58sz = len(b58u)

		outisz    = (b58sz + 3) >> 2
		binu      = make([]byte, (b58sz+3)*3)
		bytesleft = b58sz & 3

		zero = rune(encode[0])
	)

	if bytesleft > 0 {
		zmask = 0xffffffff << uint32(bytesleft*8)
	} else {
		bytesleft = 4
	}

	outi := make([]uint32, outisz)

	// расчет количества лидирующих нулей данных в кодированных данных
	for i := 0; i < b58sz && b58u[i] == zero; i++ {
		zcount++
	}

	for _, r := range b58u {
		if r > 127 {
			return nil, ErrHighBit
		}

		if decode[r] == -1 {
			return nil, fmt.Errorf("invalid base58 digit (%q)", r)
		}

		c = uint64(decode[r])

		for j := outisz - 1; j >= 0; j-- {
			t = uint64(outi[j])*58 + c
			c = (t >> 32) & 0x3f
			outi[j] = uint32(t & 0xffffffff)
		}

		if c > 0 {
			return nil, ErrOutputNumber
		}

		if outi[0]&zmask != 0 {
			return nil, ErrOutputNumber
		}
	}

	var j, cnt int

	for j, cnt = 0, 0; j < outisz; j++ {
		for mask := byte(bytesleft-1) * 8; mask <= 0x18; mask, cnt = mask-8, cnt+1 {
			binu[cnt] = byte(outi[j] >> mask)
		}

		if j == 0 {
			bytesleft = 4
		}
	}

	for n, v := range binu {
		if v > 0 {
			start := n - zcount
			if start < 0 {
				start = 0
			}

			return binu[start:cnt], nil
		}
	}

	return binu[:cnt], nil
}
