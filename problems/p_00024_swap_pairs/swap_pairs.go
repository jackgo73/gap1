package p_00024_swap_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// d-->1-->2-->3-->4
//     2-->1-->3-->4
func swapPairs(head *ListNode) *ListNode {
	var swapNextTwo func(*ListNode)
	swapNextTwo = func(head *ListNode) {
		first := head.Next
		if first == nil {
			return
		}
		second := head.Next.Next
		if second == nil {
			return
		}
		first.Next = second.Next
		second.Next = first
		head.Next = second
	}

	dummyHead := &ListNode{0, head}
	p := dummyHead
	for p.Next != nil && p.Next.Next != nil {
		swapNextTwo(p)
		p = p.Next.Next
	}
	return dummyHead.Next
}
