package Week_10

// https://leetcode-cn.com/problems/reverse-string/
// 2020-12-02

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 长度转换
func reverseString1(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
