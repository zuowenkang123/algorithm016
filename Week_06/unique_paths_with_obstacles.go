package Week_06

// https://leetcode-cn.com/problems/unique-paths-ii/
// 2020-10-09

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	// 异常返回
	if n < 1 || m < 1 {
		return 0
	}
	// 初始为障碍，返回
	if 1 == obstacleGrid[0][0] {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if 0 == i && 0 == j {
				dp[i][j] = 1
			} else if 0 == i && 0 != j {
				if 0 == obstacleGrid[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			} else if 0 != i && 0 == j {
				if 0 == obstacleGrid[i][j] {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				if 0 == obstacleGrid[i][j] {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}
