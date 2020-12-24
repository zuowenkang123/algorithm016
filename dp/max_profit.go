package dp

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// 2020-09-29

// 贪心法
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// dp 迭代模式
func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	// 第一天的情况
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	// 从后面来看，存在两种情况
	for i := 1; i < len(prices); i++ {
		// 没有不买和有卖
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 没有买和有不卖
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}
