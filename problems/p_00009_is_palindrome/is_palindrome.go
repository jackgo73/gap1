package p_00009_is_palindrome

import (
	"strconv"
)

// 121  len=3
// 1221 len=4
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	l := len(s)

	var left, right int
	if l%2 == 0 {
		mid := l/2 - 1
		left, right = palindrome(s, mid, mid+1)
	} else {
		mid := l / 2
		left, right = palindrome(s, mid, mid)
	}
	if right-left+1 == l {
		return true
	} else {
		return false
	}
}

func palindrome(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
