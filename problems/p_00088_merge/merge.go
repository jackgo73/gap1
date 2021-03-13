package p_00088_merge

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i, n := range nums2 {
		nums1[i+m] = n
	}
	sort.Slice(nums1, func(a, b int) bool {
		return nums1[a] < nums1[b]
	})
}