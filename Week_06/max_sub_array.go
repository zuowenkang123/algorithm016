package Week_06

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
		fn = max(fn+nums[i], nums[i])
		res = max(res, fn)
	}
	return res
}
