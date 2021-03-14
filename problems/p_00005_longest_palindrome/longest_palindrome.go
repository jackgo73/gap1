package p_00005_longest_palindrome

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := palindrome(s, i, i)
		left2, right2 := palindrome(s, i, i+1)
		if right2-left2 > end-start {
			start, end = left2, right2
		}
		if right1-left1 > end-start {
			start, end = left1, right1
		}
	}
	return s[start:end+1]
}

func palindrome(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}


func longestPalindrome1(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := getEdge(s, i, i)
		left2, right2 := getEdge(s, i, i+1)
		if right2-left2 > end-start {
			start, end = left2, right2
		}
		if right1-left1 > end-start {
			start, end = left1, right1
		}
	}

	return s[start: end+1]
}

func getEdge(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
