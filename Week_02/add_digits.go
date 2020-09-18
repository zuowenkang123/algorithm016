package Week_02

// https://leetcode-cn.com/problems/add-digits/
// 2020-09-18
func addDigits(num int) int {
	for num >= 10 {
		num = num/10 + num%10
	}
	return num
}

func addDigits1(num int) int {
	return (num-1)%9 + 1
}

func addDigits2(num int) int {
	if num < 10 {
		return num
	}
	// 计算当前的余数和
	next := 0
	for num != 0 {
		next = next + num%10
		num /= 10
	}
	return addDigits2(next)
}
