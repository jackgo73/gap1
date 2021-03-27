package p_00061_rotate_right

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1-2-3-4-5 4
// 2-3-4-5-1
// 0-1-2  4
// 2-0-1

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	l, r := head, head
	length := 0
	for p := head; p != nil; p = p.Next {
		length++
	}
	k = k % length

	for ; k > 0; k-- {
		r = r.Next
	}
	for r.Next != nil {
		l = l.Next
		r = r.Next
	}
	r.Next = head
	newHead := l.Next
	l.Next = nil
	return newHead
}
