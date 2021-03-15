package ms_01009_search_matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	cols, _ := len(matrix[0]), len(matrix)
	for _, row := range matrix {
		leftVal := row[0]
		rightVal := row[cols-1]

		if leftVal <= target && rightVal >= target {
			if binarySearch(row, target) != -1 {
				return true
			}
		} else {
			if leftVal > target {
				break
			}
		}
	}
	return false
}

func binarySearch(input []int, target int) int {
	left, right := 0, len(input) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if input[mid] > target {
			right = mid - 1
		} else if input[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}