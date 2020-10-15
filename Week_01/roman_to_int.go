package Week_01

// https://leetcode-cn.com/problems/roman-to-integer/
// 2020-10-15
var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	pre, res := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := romanMap[s[i]]
		if cur >= pre {
			res += cur
		} else {
			res -= cur
		}
		pre = cur
	}
	return res
}

func romanToInt1(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		// 存在并小于后面的都减去
		if i+1 < len(s) && romanMap[s[i]] < romanMap[s[i+1]] {
			res -= romanMap[s[i]]
		} else {
			res += romanMap[s[i]]
		}
	}
	return res
}
