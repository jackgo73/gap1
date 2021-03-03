package p_00100_is_same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameRecur(p, q)
}

func isSameRecur(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return isSameRecur(l.Left, r.Left) && isSameRecur(l.Right, r.Right)
}
