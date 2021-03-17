package p_00092_reverse_between

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 这时可以直接返回
	if m == n {
		return head
	}
	start := &ListNode{0, nil}
	write, cur := start, start
	// 寻找到要写入的第一个位置，也就是m之前无需反转的部分
	for i := 0; i < m; i++ {
		tmp := head.Next
		cur.Next, head.Next = head, cur.Next
		write, cur, head = cur, cur.Next, tmp
	}
	// 对 m 到 n 之间的元素执行逆序写入操作，也就是不断将新读到的节点并入到write指针之后
	for i := 0; i < n-m; i++ {
		tmp := head.Next
		write.Next, head.Next = head, write.Next
		head = tmp
	}
	// 如果写入完之后原链表还有值，要一并拼入返回值
	if head != nil {
		cur.Next = head
	}
	return start.Next
}
