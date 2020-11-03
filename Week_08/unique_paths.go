package Week_08

// https://leetcode-cn.com/problems/unique-paths/submissions/
// 2020-11-03

// 二维数组
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	//
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	// 遍历
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 一维
func uniquePaths1(m int, n int) int {
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

// 递归
func uniquePaths2(m, n int) int {
	return uniquePathsHelper(1, 1, m, n)
}

//第i行第j列到第m行第n列共有多少种路径
func uniquePathsHelper(i, j, m, n int) int {
	//边界条件的判断
	if i > m || j > n {
		return 0

	}
	if (i == m) && (j == n) {
		return 1
	}
	//从右边走有多少条路径
	right := uniquePathsHelper(i+1, j, m, n)
	//从下边走有多少条路径
	down := uniquePathsHelper(i, j+1, m, n)
	//返回总的路径
	return right + down
}

// 数学 m-1和n-1中选择m-1的个数
func uniquePaths3(m, n int) int {
	N := n + m - 2
	res := 1
	for i := 1; i < m; i++ {
		res = res * (N - (m - 1) + i) / i

	}
	return res
}
