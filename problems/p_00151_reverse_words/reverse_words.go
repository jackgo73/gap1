package p_00151_reverse_words

import "strings"

func reverseWords1(s string) string {
	if s == "" {
		return ""
	}
	s = PreProcess(s)   // 对字符串进行处理，将多余的空格去掉
	wordarry := strings.Split(s, " ")  // 切分
	i := 0
	l := len(wordarry) - 1
	for i < l {    // 交换
		wordarry[i], wordarry[l] = wordarry[l], wordarry[i]
		i++
		l--
	}
	res := strings.Join(wordarry, " ")   // 连接
	return res
}

func PreProcess(s string) string {
	l := len(s)
	var res []byte
	flag := 1   // 用于处理多个连续空格
	for l != 0 && s[l - 1] == ' ' {    // 处理字符串后面的空格
		l--
	}
	for i := 0; i < l; i++ {
		if s[i] != ' ' {
			res = append(res, s[i])
			flag = 0
		}
		if s[i] == ' ' && flag == 0 {
			res = append(res, s[i])
			flag = 1
		}
	}
	return string(res)
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	res := []string{}
	for i := len(arr) -1; i >= 0; i--  {
		str := arr[i]
		if len(str) != 0 {
			res = append(res, str)
		}
	}
	r := strings.Join(res, " ")
	return r
}