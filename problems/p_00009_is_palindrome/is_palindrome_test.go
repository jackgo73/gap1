package p_00009_is_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case0", args{1122332211}, true},
		{"case1", args{121}, true},
		{"case2", args{-121}, false},
		{"case3", args{10}, false},
		{"case4", args{-101}, false},
		{"case5", args{11223344332211}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}