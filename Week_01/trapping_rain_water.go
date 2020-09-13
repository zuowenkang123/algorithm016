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

// 动态编程
func trap1(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	ans := 0

	leftMaxArr := make([]int, len(height), len(height))
	rightMaxArr := make([]int, len(height), len(height))
	leftMaxArr[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMaxArr[i] = max(height[i], leftMaxArr[i-1])
	}
	rightMaxArr[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMaxArr[i] = max(height[i], rightMaxArr[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		ans += min(leftMaxArr[i], rightMaxArr[i]) - height[i]
	}
	return ans
}

// 栈，
func trap2(height []int) int {
	ans := 0
	stack := StackInt{}
	for i := 0; i < len(height); i++ {
		// 上升则开始计算水面积
		for false == stack.IsEmpty() && height[i] > height[stack.Top()] {
			top := stack.Top()
			// 弹出为了找到上上一个下降点进行对比
			stack, _ = stack.Pop()
			if stack.IsEmpty() {
				break
			}
			// 当前到上上一个的下降点的举例
			distance := i - stack.Top() - 1
			// 当前和上上一个下降点与上一个下降点的差值
			boundedHeight := min(height[i], height[stack.Top()]) - height[top]
			ans += distance * boundedHeight
		}
		stack = stack.Push(i)
	}
	return ans
}

// 双指针法
func trap3(height []int) int {
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
