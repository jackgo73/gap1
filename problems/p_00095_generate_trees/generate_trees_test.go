package p_00095_generate_trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test__generateTrees__O__mj(t *testing.T) {
	expect := []*TreeNode{
		{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, nil}}},
		{1, nil, &TreeNode{3, &TreeNode{2, nil, nil}, nil}},
		{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, nil},
		{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil},
	}
	res := generateTrees__O__mj(3)
	assert.EqualValues(t, expect, res)
}
