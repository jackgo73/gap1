package p_00217_contains_duplicate

func containsDuplicate(nums []int) bool {
	dupMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := dupMap[nums[i]]; ok {
			return true
		} else {
			dupMap[nums[i]] = 1
		}
	}
	return false
}
