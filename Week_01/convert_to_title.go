package Week_01

// https://leetcode-cn.com/problems/excel-sheet-column-title/

func convertToTitle(n int) string {
	arr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	reverse := ""
	for n != 0 {
		n-- // 这种1/0不对应的，先归0
		reverse += arr[n%26]
		n = n / 26
	}
	res := ""
	for i := len(reverse) - 1; i >= 0; i-- {
		res += string(reverse[i])
	}

	return res
}

// char
func convertToTitle1(n int) string {
	reverse := ""
	for n != 0 {
		n-- // 这种1/0不对应的，先归0
		reverse += string('A' + n%26)
		n = n / 26
	}
	res := ""
	for i := len(reverse) - 1; i >= 0; i-- {
		res += string(reverse[i])
	}

	return res
}

// 错误写法
func convertToTitle2(n int) string {
	arr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	reverse := ""
	for n != 0 {
		x := n % 26 // 导致0问题
		reverse += arr[x-1]
		n = n / 26 // 导致无法匹配
	}
	res := ""
	for i := len(reverse) - 1; i >= 0; i-- {
		res += string(reverse[i])
	}
	return res
}
