package p_00070_climb_stairs

// 当前 n 与 n+1 的关系
// f(n) = f(n-1)       +      f(n-2)
//        最后选择爬1个         最后选择爬2个
// f(0) = 1
// f(1) = 1
// f(2) = 2
// f(3) = 3
// f(4) = 5

func climbStairs(n int) int {
	trans := make([]int, n+1, n+1)
	trans[0], trans[1] = 1, 1

	for i := 2; i < n+1; i++ {
		trans[i] = trans[i-1] + trans[i-2]
	}

	return trans[n]
}
func climbStairs1(n int) int {
	trans := make([]int, n+1)
	trans[0] = 1
	trans[1] = 1
	for i := 2; i < n+1; i++ {
		trans[i] = trans[i-1] + trans[i-2]
	}
	return trans[n]
}
