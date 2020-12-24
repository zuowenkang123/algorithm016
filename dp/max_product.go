package dp

// https://leetcode-cn.com/problems/maximum-product-subarray/description/
// 2020-10-11
func maxProduct(nums []int) int {
	maxF, minF, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		res = max(maxF, res)
	}
	return res
}
