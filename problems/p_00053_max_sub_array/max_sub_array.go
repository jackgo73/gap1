package p_00053_max_sub_array

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	// -1<<63 math.MinInt64
	max := math.MinInt64
	// 1<<63-1 math.MaxInt64
	//max1 := math.MaxInt64
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
