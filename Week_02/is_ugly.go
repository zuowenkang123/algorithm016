package Week_02

// https://leetcode-cn.com/problems/ugly-number/
// 2020-12-05

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}
	return true
}
