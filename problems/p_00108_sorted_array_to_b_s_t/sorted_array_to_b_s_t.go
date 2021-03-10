package p_00108_sorted_array_to_b_s_t

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return creator(nums, 0, len(nums)-1)
}

func creator(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	node := &TreeNode{nums[mid], nil, nil}
	node.Left = creator(nums, left, mid-1)
	node.Right = creator(nums, mid+1, right)
	return node
}
