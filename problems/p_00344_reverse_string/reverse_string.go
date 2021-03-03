package p_00344_reverse_string

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for ; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}
