package p_00004_find_median_sorted_arrays

func findMedianSortedArrays__n__mj(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	lt := l1 + l2
	left, right := -1, -1
	i1, i2 := 0, 0

	for i := 0; i <= lt/2; i++ {
		left = right
		if i1 < l1 && (i2 >= l2 || nums1[i1] < nums2[i2]) {
			right = nums1[i1]
			i1 += 1
		} else {
			right = nums2[i2]
			i2 += 1
		}
	}
	if lt&1 == 0 {
		return float64(left+right) / 2.0
	} else {
		return float64(right)
	}
}
