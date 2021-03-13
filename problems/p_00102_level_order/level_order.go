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

func levelOrder2(root *TreeNode) [][]int {
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


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	currLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}

	for len(currLevel) != 0 {
		currLevelVal := []int{}
		for i := 0; i < len(currLevel); i++ {
			node := currLevel[i]
			currLevelVal = append(currLevelVal, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, currLevelVal)
		currLevel = nextLevel
	}
	return res
}