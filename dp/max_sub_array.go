package dp

// https://leetcode-cn.com/problems/maximum-subarray
// 2020-10-08

// dp 递推
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := -1 << 31
	fn := -1
	for i := 0; i < len(nums); i++ {
		// 当前结尾最大
		fn = max(fn+nums[i], nums[i])
		// 整体最大
		res = max(res, fn)
	}
	return res
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		// 当前结尾最大
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		// 整体最大
		res = max(res, dp[i])
	}
	return res
}
