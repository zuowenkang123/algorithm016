package Week_01

// https://leetcode-cn.com/problems/trapping-rain-water/submissions/
// 2020-09-13 22:24:33

// 暴力法
func trap(height []int) int {
	ans := 0
	for i := 0; i < len(height); i++ {
		left_max := 0
		right_max := 0
		for j := i; j >= 0; j-- {
			left_max = max(height[j], left_max)
		}
		for j := i; j < len(height); j++ {
			right_max = max(height[j], right_max)
		}
		ans += min(left_max, right_max) - height[i]
	}
	return ans
}

// 双指针法
func trap1(height []int) int {
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
