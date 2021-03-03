package p_00094_inorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 *        1
 *     3     5
 *   7   6     9
 *  [1, 3, 7, 6, 5, 9]
 */
func Test__inorderTraversal__O__mj(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{3, &TreeNode{Val: 7}, &TreeNode{Val: 6}},
		&TreeNode{5, nil, &TreeNode{Val: 9}},
	}
	res := inorderTraversal__O__mj(root)
	assert.EqualValues(t, []int{7, 3, 6, 1, 5, 9}, res)
}
