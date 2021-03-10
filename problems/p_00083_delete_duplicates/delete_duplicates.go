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
