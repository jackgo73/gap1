package ms_00207_get_intersection_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	listMap := map[*ListNode]int{}
	for headA != nil {
		listMap[headA] = 1
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := listMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
