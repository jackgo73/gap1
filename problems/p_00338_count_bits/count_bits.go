package p_00338_count_bits

func onesCount1(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(num int) []int {
	bits := make([]int, num+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

func onesCount(x int) (ones int) {
	for ; x > 0; x = x >> 1 {
		ones = ones + (x & 1)
	}
	return
}
