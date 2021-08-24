package base58

const (
	btcSymbols    = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	flikrSymbols  = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	rippleSymbols = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
)

type alphabet struct {
	decode [128]int8
	encode [58]byte
}

func BTCAlphabet() Alphabet {
	return newAlphabet(btcSymbols)
}

func FlikrAlphabet() Alphabet {
	return newAlphabet(flikrSymbols)
}

func RippleAlphabet() Alphabet {
	return newAlphabet(rippleSymbols)
}

func newAlphabet(symb string) Alphabet {
	abt := &alphabet{}
	
	copy(abt.encode[:], symb)
	
	for i := 0; i < len(abt.decode); i++ {
		abt.decode[i] = -1
	}
	
	for i := 0; i < len(abt.decode); i++ {
		abt.decode[abt.encode[i]] = int8(i)
	}
	
	return abt
}

func (a alphabet) EncodeSymbols() [58]byte {
	return a.encode
}

func (a alphabet) DecodeSymbols() [128]int8 {
	return a.decode
}
