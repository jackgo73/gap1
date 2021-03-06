package p_00392_is_subsequence

// s = "abc", t = "ahbgdc"
// s = "axc", t = "ahbgdc"
func isSubsequence(s string, t string) bool {
	sp, tp := 0, 0
	for sp < len(s) && tp < len(t) {
		if s[sp] == t[tp] {
			sp++
		}
		tp++
	}
	res := sp == len(s)
	return res
}
