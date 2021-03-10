

# index
## 1 回文

所有回文问题统一使用函数
```go
func palindrome(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
```

[p_00005_longest_palindrome](problems/p_00005_longest_palindrome)
[p_00009_is_palindrome](problems/p_00009_is_palindrome)

