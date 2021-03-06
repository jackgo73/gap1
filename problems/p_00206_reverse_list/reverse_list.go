package p_00206_reverse_list

type ListNode struct {
	Val int
	Next *ListNode
}

// 1->2->3->4->5->NULL
// 5->4->3->2->1->NULL
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}