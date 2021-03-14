package mj

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := test()
			fmt.Println(s)
		})
	}
}
