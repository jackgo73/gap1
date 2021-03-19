package p_00147_insertion_sort_list

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	arr := []int{}
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})
	hdr := &ListNode{}
	p := hdr
	for _, v := range arr {
		p.Next = &ListNode{v, nil}
		p = p.Next
	}
	p.Next = nil
	return hdr.Next
}

func insertionSortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
