package p_00053_max_sub_array

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if sum >= 0 {
			sum += num
		} else {
			sum = num
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
