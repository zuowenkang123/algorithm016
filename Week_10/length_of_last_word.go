package Week_10

// https://leetcode-cn.com/problems/length-of-last-word/
// 2020-12-02 16:28:35

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res == 0 {
				continue
			}
			break
		} else {
			res++
		}
	}
	return res
}
