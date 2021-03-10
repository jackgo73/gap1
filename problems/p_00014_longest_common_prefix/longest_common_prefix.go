package p_00014_longest_common_prefix

// - - - - - - - - - i
// | f l o w e r
// | f l o w
// | f l i g h t
// |
// |
// j
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
