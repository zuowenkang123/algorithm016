package Week_01

// 暴力法，最后处理进位的问题
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		// 只有9才会为0，因为是加1
		if digits[i] != 0 {
			return digits
		}
	}

	res := make([]int, 0, len(digits)+1)
	res = append(res, 1)
	res = append(res, digits...)
	return res
}

// 进位可见，常用方法
func plusOne1(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		carry = (digits[i] + carry) / 10
		digits[i] = (digits[i] + 1) % 10
		if carry == 0 {
			break
		}
	}

	if carry == 0 {
		return digits
	}
	res := make([]int, 0, len(digits))
	res = append(res, carry)
	for _, val := range digits {
		res = append(res, val)
	}
	return res
}
