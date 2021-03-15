package p_00142_detect_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	memMap := map[*ListNode]bool{}
	for head != nil {

		if _, ok := memMap[head]; ok {
			return head
		} else {
			memMap[head] = true
		}
		head = head.Next
	}
	return nil
}
