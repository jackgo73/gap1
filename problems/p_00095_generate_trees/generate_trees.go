package p_00095_generate_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees__O__mj(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
    var generator func(start, end int) []*TreeNode
	generator = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		allTrees := []*TreeNode{}
		// 枚举可行根节点
		for i := start; i <= end; i++ {
			// 获得所有可行的左子树集合
			leftTrees := generator(start, i-1)
			// 获得所有可行的右子树集合
			rightTrees := generator(i+1, end)
			// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					currTree := &TreeNode{i, nil, nil}
					currTree.Left = left
					currTree.Right = right
					allTrees = append(allTrees, currTree)
				}
			}
		}
		return allTrees
	}
	return generator(1, n)
}
