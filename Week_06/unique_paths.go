package Week_06

// https://leetcode-cn.com/problems/unique-paths/
// 2020-10-09

// 定义最左列，然后逐步累加至右侧
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

// 先两边，再递推
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if 0 == i || 0 == j {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
