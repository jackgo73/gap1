package p_00002_add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHdr := &ListNode{0, nil}
	p := dummyHdr
	sum := 0
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum = n1 + n2 + carry
		if sum >= 10 {
			p.Next = &ListNode{sum % 10, nil}
			carry = 1
		} else {
			p.Next = &ListNode{sum, nil}
			carry = 0
		}
		p = p.Next
	}
	if carry == 1 {
		p.Next = &ListNode{Val: 1}
	}
	return dummyHdr.Next
}
