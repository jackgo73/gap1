package p_00028_str_str

func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	if l2 == 0 {
		return 0
	}
	if l2 > l1 || l1 == 0{
		return -1
	}

	for i := 0; i < l1; i++ {
		j, k := i, 0
		for ; ; j, k = j+1, k+1 {
			if j >= l1 || k >= l2 || haystack[j] != needle[k] {
				break
			}
		}
		if k == l2 {
			return j - k
		}
	}
	return -1
}
