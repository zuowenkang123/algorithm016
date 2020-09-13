package Week_01

// https://leetcode-cn.com/problems/trapping-rain-water/submissions/
// 2020-09-13 22:24:33

// 双指针法
func trap(height []int) int {
	left := 0
	right := len(height) - 1
	left_max := 0
	right_max := 0
	ans := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= left_max {
				left_max = height[left]
			} else {
				ans += left_max - height[left]
			}
			left++
		} else {
			if height[right] >= right_max {
				right_max = height[right]
			} else {
				ans += right_max - height[right]
			}
			right--
		}
	}
	return ans
}
