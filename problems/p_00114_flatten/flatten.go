package p_00114_flatten

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	// stop
	if root == nil {
		return
	}
	// [pre]/in/post

	// recur
	flatten(root.Left)
	flatten(root.Right)

	// pre/in/[post]
	left := root.Left
	right := root.Right

	// 左子 挂到 右面
	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}
