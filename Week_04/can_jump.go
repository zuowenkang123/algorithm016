package Week_04

// https://leetcode-cn.com/problems/jump-game/
// 2020-10-11

// 贪心
// 从最后开始，如果满足则更新可到达标记
func canJump(nums []int) bool {
	left := len(nums) - 1
	for i := left - 1; i >= 0; i-- {
		if nums[i]+i >= left {
			left = i
		}
	}
	return left == 0
}

// 贪心
// 从最后开始，如果满足则更新可到达标记
func canJump1(nums []int) bool {
	rightMost := 0
	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			rightMost = max(rightMost, i+nums[i])
			if rightMost >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

// dp 递推每一个可能
func canJump3(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 1; i < len(nums); i++ {
		flag := false
		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				flag = true
				break
			}
		}
		dp[i] = flag
	}

	return dp[len(nums)-1]
}

// 错误写法，只能追溯前一个
func canJump2(nums []int) bool {
	numsLen := len(nums)
	if len(nums) <= 1 {
		return false
	}

	if canJump(nums[:numsLen-1]) && nums[numsLen-2] >= 1 {
		return true
	}
	return false
}
