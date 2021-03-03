package p_00167_two_sum

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left := i + 1
		right := len(numbers) - 1
		for left <= right {
			mid := left + (right-left)/2
			if target - numbers[i] == numbers[mid] {
				res := []int{i + 1, mid + 1}
				return res
			} else if target - numbers[i] > numbers[mid] {
				left = mid + 1
			} else if target - numbers[i] < numbers[mid] {
				right = mid - 1
			}
		}
	}
	return []int{-1, -1}
}
