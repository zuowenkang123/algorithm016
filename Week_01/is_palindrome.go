package Week_01

import "strconv"

// https://leetcode-cn.com/problems/palindrome-number/
// 2020-10-15

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)
	left := 0
	right := len(xStr) - 1
	for left <= right {
		if xStr[left] != xStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	y := reverseInt(x)
	if x == y {
		return true
	}
	return false
}
