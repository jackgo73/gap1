package p_00050_my_pow

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var pow float64 = 1
	for n > 0 {
		//如果n为奇数，则先乘以x
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n = n >> 1
	}
	return pow
}
