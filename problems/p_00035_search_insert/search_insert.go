package p_00035_search_insert

func searchInsert(nums []int, target int) int {

	left, right := 0, len(nums)-1
	res := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
			res = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
			res = mid
		}
	}
	return res
}
