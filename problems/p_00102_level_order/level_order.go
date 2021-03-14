package p_00102_level_order

import "fmt"

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

func levelOrder1(root *TreeNode) [][]int {
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

	s := []int{1,2,3,4,5,6}
	fmt.Println(s[0:2])
	fmt.Println(s[1:2])
	fmt.Println(s[:6])
	fmt.Println(s[1:])


	res := [][]int{}
	currLevel := []*TreeNode{root}


	for len(currLevel) != 0 {
		currLevelVal := []int{}
		nextLevel := []*TreeNode{}
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


func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	currLevel := []*TreeNode{root}
	// level
	for len(currLevel) > 0 {
		// level inside
		nextLevel := []*TreeNode{}
		currVals := []int{}
		for i := 0; i < len(currLevel); i++ {
			node := currLevel[i]
			currVals = append(currVals, node.Val)

			// fill next level
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, currVals)
		currLevel = nextLevel
	}
	return res
}