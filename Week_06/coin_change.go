package Week_06

// https://leetcode-cn.com/problems/coin-change/submissions/
// 2020-10-11
// https://leetcode-cn.com/problems/coin-change/solution/javadi-gui-ji-yi-hua-sou-suo-dong-tai-gui-hua-by-s/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				if dp[i] == -1 || dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for k := range dp {
		dp[k] = amount + 1
	}
	dp[0] = 0
	for i := 0; i <= amount; i++ {
		// 内层 for 在求所有子问题 + 1 的最小值
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
