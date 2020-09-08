package Week_01

// 暴力法，最后处理进位的问题
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		carry = (digits[i] + carry) / 10
		digits[i] = (digits[i] + 1) % 10
		if carry == 0 {
			break
		}
	}
	res := make([]int, 0, len(digits))
	if carry == 1 {
		res = append(res, carry)
	}
	for _, val := range digits {
		res = append(res, val)
	}
	return res
}
