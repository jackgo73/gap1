package p_00083_delete_duplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := head, head
	for ; right != nil; right = right.Next {
		if left.Val != right.Val {
			left.Next = right
			left = left.Next
		}
	}
	left.Next = nil
	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}

	}
	return head
}
