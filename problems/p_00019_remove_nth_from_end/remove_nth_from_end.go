package p_00019_remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3 4 5 6
// 2
// 1 2 3 4   6
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p_head, p_follow := head, head

	for ; n > 0; n-- {
		p_head = p_head.Next
	}

	if p_head == nil {
		return head.Next
	}

	for p_head != nil && p_head.Next != nil {
		p_head = p_head.Next
		p_follow = p_follow.Next
	}

	p_follow.Next = p_follow.Next.Next
	return head
}
