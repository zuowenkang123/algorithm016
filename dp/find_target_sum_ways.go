package dp

// https://leetcode-cn.com/problems/target-sum/
func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if (sum+S)%2 == 1 {
		return 0
	}

	if S > sum {
		return 0
	}
	x := (sum + S) / 2
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, x+1)
	}
	dp[0][0] = 1

	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= x; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)][x]
}

// todo 待调试
func findTargetSumWays1(nums []int, S int) int {
	// 和
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// 奇数不要
	if (sum+S)%2 == 1 {
		return 0
	}
	// 差大于和不要
	if S > sum {
		return 0
	}
	// 01背包目标
	x := (sum + S) / 2

	dp := make([]int, len(nums)+1)
	dp[0] = 1

	for _, num := range nums {
		for i := x; i > num; i-- {
			dp[i] = dp[i-num]
		}
	}

	return dp[x]
}
