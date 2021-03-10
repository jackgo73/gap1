package p_00027_remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != val {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
