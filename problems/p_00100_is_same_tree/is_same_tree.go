package p_00100_is_same_tree
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	return isSameRecur(p, q)
//}
//
//func isSameRecur(l *TreeNode, r *TreeNode) bool {
//	if l == nil && r == nil {
//		return true
//	}
//
//	if l == nil || r == nil {
//		return false
//	}
//
//	if l.Val != r.Val {
//		return false
//	}
//
//	return isSameRecur(l.Left, r.Left) && isSameRecur(l.Right, r.Right)
//}
//
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	// end
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	// pre
	if q.Val != p.Val {
		return false
	}
	// rescur
	l := isSameTree(p.Left, q.Left)
	r := isSameTree(p.Right, q.Right)
	return l && r
}
