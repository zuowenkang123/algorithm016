package Week_04

// https://leetcode-cn.com/problems/jump-game-ii/
// 2020-10-11 16:44:45

// 向前更新
func jump(nums []int) int {
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
