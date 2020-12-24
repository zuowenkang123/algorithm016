package dp

// https://leetcode-cn.com/problems/triangle/solution/golangban-triangle-by-studyzy/
// 2020-10-11

// 自底而上
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

// 一维方式，最小归并到行首
func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
