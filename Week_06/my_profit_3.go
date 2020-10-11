package Week_06

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/dong-tai-gui-hua-by-ling-chen-12dian/
// 2020-10-12
func maxProfit_3(prices []int) int {
	pLen := len(prices)
	if pLen == 0 {
		return 0
	}

	result := 0
	profit := make([][3][2]int, pLen)

	profit[0][0][0], profit[0][0][1] = 0, -prices[0]
	profit[0][1][0], profit[0][1][1] = 0, -prices[0]
	profit[0][2][0], profit[0][2][1] = 0, -prices[0]

	//注意 买 卖 都进行一次算一次操作 k + 1,单独 买入 不算完成一次操作
	for i := 1; i < pLen; i++ {
		//第 i 天 0 次交易 没有股票最大利润 = 第 i-1 天 0 次交易 没有股票最大利润
		profit[i][0][0] = profit[i-1][0][0]
		//第 i 天 0 次交易 有股票最大利润 = max(第 i-1 天 0 次交易 有股票最大利润 , 第 i-1 天 0 次交易 无股票最大利润 - 当天股票价格prices[i]（买入）)
		profit[i][0][1] = max(profit[i-1][0][1], profit[i-1][0][0]-prices[i])

		//# 第 i 天 1 次交易 无股票最大利润 = max(第 i-1 天 1次交易 无股票最大利润 , 第 i-1 天 0 次交易 有股票最大利润 + 当天股票价格prices[i]（卖出）)
		profit[i][1][0] = max(profit[i-1][1][0], profit[i-1][0][1]+prices[i])
		// # 第 i 天 1 次交易 有股票最大利润 = max(第 i-1 天 1 次交易 有股票最大利润 , 第 i-1 天 1 次交易 无股票最大利润 - 当天股票价格prices[i]（买入）)
		profit[i][1][1] = max(profit[i-1][1][1], profit[i-1][1][0]-prices[i])

		//# 第 i 天 2 次交易 无股票最大利润 = max(第 i-1 天 2次交易 无股票最大利润 , 第 i-1 天 1 次交易 有股票最大利润 + 当天股票价格prices[i]（卖出）)
		profit[i][2][0] = max(profit[i-1][2][0], profit[i-1][1][1]+prices[i])

	}
	result = max(profit[pLen-1][0][0], max(profit[pLen-1][1][0], profit[pLen-1][2][0]))
	return result
}
