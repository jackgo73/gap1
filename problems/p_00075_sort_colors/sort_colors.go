package p_00075_sort_colors

import "sort"

func sortColors(nums []int)  {
	sort.Slice(nums, func(a,b int) bool {
		return nums[a] < nums[b]
	})
}