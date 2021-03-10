package p_00026_remove_duplicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{"case1", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, 5, []int{0, 1, 2, 3, 4}},
		{"case2", args{[]int{1, 1, 2}}, 2, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				for i := 0; i < tt.want; i++ {
					assert.Equal(t, tt.args.nums[i], tt.want1[i])
				}
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
