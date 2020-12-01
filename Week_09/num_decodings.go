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
		// 10/20的情况
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			} else {
				cur = pre
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') { // 1x/2x的情况
			cur += pre
		}
		pre = tmp // 个位的情况
	}
	return cur
}

func numDecodings1(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < len(s); i++ {
		// 10/20的情况
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			} else {
				dp[i+1] = dp[i-1]
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') { // 1x/2x的情况
			dp[i+1] = dp[i] + dp[i-1]
		} else {
			dp[i+1] = dp[i] // 个位的情况
		}
	}
	return dp[len(s)]
}
