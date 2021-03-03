package p_00509_fib

// 2^n
func fib__O2n(n int) int {
	var rescur func(n int) int
	rescur = func(n int) int {
		if n == 1 || n == 2 {
			return 1
		}
		return rescur(n-1) + rescur(n-2)
	}
	return rescur(n)
}

// n
func fib__On__memo(n int) int {
	if n < 1 {
		return 0
	}
	memo := make([]int, n+1)
	var rescur func(n int) int
	rescur = func(n int) int {
		if n == 1 || n == 2 {
			return 1
		}
		if memo[n] != 0 {
			return memo[n]
		}
		memo[n] = rescur(n-1) + rescur(n-2)
		return memo[n]
	}
	return rescur(n)
}

// n
func fib__On__dp(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
