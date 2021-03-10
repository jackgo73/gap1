package p_00219_contains_nearby_duplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	dupMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := dupMap[nums[i]]; ok {
			if i - j <= k {
				return true
			} else {
				dupMap[nums[i]] = i
			}
		} else {
			dupMap[nums[i]] = i
		}
	}
	return false
}
