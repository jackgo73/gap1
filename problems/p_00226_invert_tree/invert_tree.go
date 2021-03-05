package p_00226_invert_tree

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// st
	if root == nil {
		return nil
	}

	// fr
	root.Left, root.Right = root.Right, root.Left

	// tr
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
