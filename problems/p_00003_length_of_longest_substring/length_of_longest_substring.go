package p_00003_length_of_longest_substring

// a b c a b c b b
// |     |
// a b c b a c b b
// |     |
//   |   |
//     | |
func lengthOfLongestSubstring(s string) int {
	winMap := make(map[byte]int)
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		winMap[c]++
		for winMap[c] > 1 {
			d := s[left]
			left++
			winMap[d]--
		}
		res = maxInt(res, right-left)
	}
	return res
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//
//func lengthOfLongestSubstring(s string) int {
//	strMap := map[byte]int{
//		s[0] : 0,
//	}
//	left, right := 0, 1
//	max := -1<<63
//
//	// move right
//	for right < len(s) {
//		c := s[right]
//		if i, ok := strMap[c]; !ok {
//			strMap[c] = right
//		} else {
//			//move left
//			left = i + 1
//			strMap[c] = right
//		}
//		if l := right - left; l > max {
//			max = l
//		}
//		right++
//	}
//	return max + 1
//}

