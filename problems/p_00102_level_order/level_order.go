package p_00102_level_order

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//   3
//  / \
// 9  20
//   /  \
//  15   7

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	curLevel := []*TreeNode{root}
	for i := 0; len(curLevel) > 0; i++ {
		ret = append(ret, []int{})
		nextLevel := []*TreeNode{}
		for j := 0; j < len(curLevel); j++ {
			node := curLevel[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		curLevel = nextLevel
	}
	return ret
}
