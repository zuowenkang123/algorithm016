package dp

// http://www.voidcn.com/article/p-expzbqhu-bsd.html
/*
给定一个有n个正整数的数组A和一个整数sum,求选择数组A中部分数字和为sum的方案数。
当两种选取方案有一个数字的下标不一样,我们就认为是不同的组成方案。
输入描述:
输入为两行:
 第一行为两个正整数n(1 ≤ n ≤ 1000)，sum(1 ≤ sum ≤ 1000)
 第二行为n个正整数A[i](32位整数)，以空格隔开。
输出描述:
输出所求的方案数
*/

// 本体类似can_partition

func NSum(nums []int, sum int) int {
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	for j := 1; j <= sum; j++ {
		dp[0][j] = 0
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= sum; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}
