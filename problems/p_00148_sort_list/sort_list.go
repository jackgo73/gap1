package p_00148_sort_list

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})
	dHead := &ListNode{}
	p := dHead
	for _, n := range arr {
		p.Next = &ListNode{n, nil}
		p = p.Next
	}
	p.Next = nil
	return dHead.Next
}
