package dp

// https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	// l表示距离
	for l := 0; l < n; l++ {
		// 从头开始，直到最后
		for i := 0; i+l < n; i++ {
			j := i + l
			// 自己
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 { // 成对
				if s[i] == s[j] {
					dp[i][j] = true
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && l+1 > len(ans) { // 求最大
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}

// 最长回文子串
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
