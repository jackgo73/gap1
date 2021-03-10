package p_00058_length_of_last_word

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
	return len(split[len(split)-1])
}