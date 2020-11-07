package Week_08

// https://leetcode-cn.com/problems/longest-increasing-subsequence/
// 2020-11-05

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		// 默认当前
		dp[i] = 1
		// 对比之前
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				// 累加，对比之前所有的j，找出最大的
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	return res
}
