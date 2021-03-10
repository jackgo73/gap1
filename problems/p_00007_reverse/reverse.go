package p_00007_reverse

import "math"

func reverse(x int) int {
	res := 0
	for x != 0 {
		i := x % 10
		x /= 10
		res = res*10 + i
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}
