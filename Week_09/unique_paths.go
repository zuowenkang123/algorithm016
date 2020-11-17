package Week_09

// https://leetcode-cn.com/problems/unique-paths/
// 2020-11-15 13:51:35

// 二维递推解法
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	// 创建存储
	dp := make([][]int, m)
	// 分配子存储
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始边
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	// 初始边
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 一维递推解法
func uniquePaths1(m int, n int) int {
	// 创建存储
	dp := make([]int, n)
	// 分配子存储
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 每一步都是新的路径
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
