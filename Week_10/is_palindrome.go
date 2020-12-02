package Week_10

import "strings"

// https://leetcode-cn.com/problems/valid-palindrome/
// 2020-12-02
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)
	for left <= right {
		if !isValid(s[left]) {
			left++
			continue
		}
		if !isValid(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isValid(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
