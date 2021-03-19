package p_00086_partition

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	leftDummyHdr := &ListNode{}
	l := leftDummyHdr
	rightDummyHdr := &ListNode{}
	r := rightDummyHdr

	p := head
	for p != nil {
		if p.Val < x {
			l.Next = p
			l = l.Next
		} else {
			r.Next = p
			r = r.Next
		}
		p = p.Next
	}
	r.Next = nil
	l.Next = rightDummyHdr.Next
	return leftDummyHdr.Next
}
