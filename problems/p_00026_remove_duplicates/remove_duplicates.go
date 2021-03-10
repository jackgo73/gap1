package p_00026_remove_duplicates

// 0 0 1 1 1 2 2 3 3 4
func removeDuplicates(nums []int) (cnt int) {
	if len(nums) == 0 {
		return 0
	}
	for left, right := 0, 1; right < len(nums); right++ {
		if nums[right] > nums[left] {
			left += 1
			nums[left] = nums[right]
			cnt += 1
		}
	}
	return cnt + 1
}
