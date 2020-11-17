package Week_09

// https://leetcode-cn.com/problems/decode-ways/
// 2020-11-17

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	cur, pre := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			} else {
				cur = pre
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			cur += pre
		}
		pre = tmp
	}
	return cur
}
