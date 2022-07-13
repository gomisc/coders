package base58

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func Test_encoder_Encode(t *testing.T) {
	type fields struct {
		alphabet Alphabet
	}

	type args struct {
		in []byte
	}

	idBytes, _ := hex.DecodeString("eede0fac3c3543a293ef03b959a6d5ac")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "positive btc",
			fields: fields{alphabet: newAlphabet(btcSymbols)},
			args:   args{in: idBytes},
			want:   "WVnnPdFu8qiwEiyYAmhxCK",
		},
		{
			name:   "positive flikr",
			fields: fields{alphabet: newAlphabet(flikrSymbols)},
			args:   args{in: idBytes},
			want:   "vuMMoCfU8QHWeHYxaLGXcj",
		},
		{
			name:   "positive ripple",
			fields: fields{alphabet: newAlphabet(rippleSymbols)},
			args:   args{in: idBytes},
			want:   "WV88PdEu3q5AN5yYwm6xUK",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			e := &encoder{
				alphabet: tt.fields.alphabet,
			}

			if got := e.Encode(tt.args.in); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encoder_Decode(t *testing.T) {
	type fields struct {
		alphabet Alphabet
	}

	type args struct {
		in string
	}

	idBytes, _ := hex.DecodeString("eede0fac3c3543a293ef03b959a6d5ac")

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "positive btc",
			fields: fields{alphabet: newAlphabet(btcSymbols)},
			args:   args{in: "WVnnPdFu8qiwEiyYAmhxCK"},
			want:   idBytes,
		},
		{
			name:   "positive flikr",
			fields: fields{alphabet: newAlphabet(flikrSymbols)},
			args:   args{in: "vuMMoCfU8QHWeHYxaLGXcj"},
			want:   idBytes,
		},
		{
			name:   "positive ripple",
			fields: fields{alphabet: newAlphabet(rippleSymbols)},
			args:   args{in: "WV88PdEu3q5AN5yYwm6xUK"},
			want:   idBytes,
		},
		{
			name:    "negative empty arg",
			fields:  fields{alphabet: newAlphabet(btcSymbols)},
			args:    args{in: ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			e := &encoder{
				alphabet: tt.fields.alphabet,
			}
			got, err := e.Decode(tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
