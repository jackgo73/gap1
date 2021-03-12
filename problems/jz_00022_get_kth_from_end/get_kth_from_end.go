package jz_00022_get_kth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	left, right := head, head

	for i := 0; i < k-1; i++ {
		right = right.Next
		if right == nil {
			return nil
		}
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	return left
}
