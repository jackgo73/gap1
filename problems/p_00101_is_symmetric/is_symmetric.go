package p_00101_is_symmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	l := check(p.Left, q.Right)
	r := check(p.Right, q.Left)
	return l && r

}



func isSymmetric1(root *TreeNode) bool {
	return isSame(root,root)
}

func isSame(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	l := isSame(left.Left, right.Right)
	r := isSame(left.Right, right.Left)
	return l && r
}