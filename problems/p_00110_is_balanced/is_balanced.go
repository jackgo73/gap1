package p_00110_is_balanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := height(root.Left)
	r := height(root.Right)
	if l-r > 1 || l-r < -1 {
		return false
	}
	if !isBalanced(root.Left) {
		return false
	}
	if !isBalanced(root.Right) {
		return false
	}
	return true
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := height(node.Left)
	r := height(node.Right)
	if l >= r {
		return l + 1
	} else {
		return r + 1
	}
}
