package p_00226_invert_tree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"case1",
			args{root: &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}},
			           &TreeNode{4, &TreeNode{7, &TreeNode{9, nil, nil}, &TreeNode{6, nil, nil}}, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
