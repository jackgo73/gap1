package p_00020_is_valid

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{s: "()"}, true},
		{"case2", args{s: "()[]{}"}, true},
		{"case3", args{s: "(]"}, false},
		{"case4", args{s: "([)]"}, false},
		{"case5", args{s: "{[]}"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
