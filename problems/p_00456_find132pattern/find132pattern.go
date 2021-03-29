package p_00456_find132pattern

func find132pattern(nums []int) bool {
	nums_i := nums[0]

	for j := 1; j < len(nums) - 1; j++ {
		for k := len(nums) - 1; k > j; k-- {
			if nums[j] > nums_i && nums[j] > nums[k] && nums[k] > nums_i {
				return true
			}
		}
		if nums_i > nums[j] {
			nums_i = nums[j]
		}
	}
	return false
}