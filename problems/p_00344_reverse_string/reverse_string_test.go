package p_00344_reverse_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{"case1", args{[]byte{'h', 'e', 'l', 'l', 'o'}}, args{[]byte{'o', 'l', 'l', 'e', 'h'}}},
		{"case2", args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}}, args{[]byte{'h', 'a', 'n', 'n', 'a', 'H'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			assert.EqualValues(t, tt.args.s, tt.want.s)
		})
	}
}
