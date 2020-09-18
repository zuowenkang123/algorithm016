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
