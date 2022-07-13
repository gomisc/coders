package uuid

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestUUIDv5(t *testing.T) {
	type args struct {
		ns  uuid
		arg string
	}

	bid, _ := hex.DecodeString("6de4b5af077357dfbd6bcee55ea8e595")

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "positive",
			args: args{
				ns:  Zero,
				arg: "тестоваястрока",
			},
			want: bid,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := UUIDv5(tt.args.ns, tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUIDv5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrUUIDv5(t *testing.T) {
	type args struct {
		ns  uuid
		arg string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive",
			args: args{
				ns:  Zero,
				arg: "тестоваястрока",
			},
			want: "6de4b5af-0773-57df-bd6b-cee55ea8e595",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := StrUUIDv5(tt.args.ns, tt.args.arg); got != tt.want {
				t.Errorf("StrUUIDv5() = %v, want %v", got, tt.want)
			}
		})
	}
}
