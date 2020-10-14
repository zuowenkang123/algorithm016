package Week_01

// https://leetcode-cn.com/problems/container-with-most-water/
// 2020-09-07

// 双向指针，左右移动得到的底是相同的，因此小的寻求更大的机会
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	begin := 0
	end := len(height) - 1
	res := 0
	for begin < end {
		res = max(res, (end-begin)*min(height[begin], height[end]))
		// 左右移动得到的底是相同的，因此小的寻求更大的机会
		if height[begin] > height[end] {
			end--
		} else {
			begin++
		}
	}
	return res
}

// 暴力法
func maxArea1(height []int) int {
	if len(height) < 2 {
		return 0
	}
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			res = max(res, (j-i)*min(height[i], height[j]))
		}
	}
	return res
}
