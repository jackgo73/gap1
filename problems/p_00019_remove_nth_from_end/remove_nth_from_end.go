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



func removeNthFromEnd1(head *ListNode, n int) *ListNode {

	dummyheader := head
	left, right := head, head

	for n > 0 {
		right = right.Next
		n--
	}

	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next

	return dummyheader
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	l, r := head, head

	for ; n > 0; n-- {
		r = r.Next
	}

	if r == nil {
		return head.Next
	}

	for r.Next != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next

	return head
}