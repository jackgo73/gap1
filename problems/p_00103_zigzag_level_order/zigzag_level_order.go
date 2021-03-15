package p_00103_zigzag_level_order

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode)  [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	currentLevel := []*TreeNode{root}
	for level := 0; len(currentLevel) > 0; level++ {
		vals := []int{}
		q := currentLevel
		currentLevel = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				currentLevel = append(currentLevel, node.Left)
			}
			if node.Right != nil {
				currentLevel = append(currentLevel, node.Right)
			}
		}

		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		res = append(res, vals)
	}
	return res
}
