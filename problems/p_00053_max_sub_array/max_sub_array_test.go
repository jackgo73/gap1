package p_00053_max_sub_array

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{"case1", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"case2", args{[]int{-1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxSubArray(tt.args.nums); gotMax != tt.wantMax {
				t.Errorf("maxSubArray() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
