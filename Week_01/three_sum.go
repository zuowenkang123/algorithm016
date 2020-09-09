package Week_01

import (
	"sort"
)

// https://leetcode-cn.com/problems/3sum/
// 2020-09-08
// 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resArr := make([][]int, 0)
	if len(nums) < 3 {
		return resArr
	}
	for i := 0; i < len(nums); i++ {
		// 重复判断，当前位置对比之前。还可以map
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		head := i + 1
		tail := len(nums) - 1

		for head < tail {
			if nums[head]+nums[tail]+nums[i] == 0 {
				// 添加
				resArr = append(resArr, []int{nums[head], nums[tail], nums[i]})
				// 处理左重复逻辑，右边的会被兼容
				for head < tail && nums[head] == nums[head+1] {
					head++
				}
				head++
				tail--
			} else if nums[head]+nums[tail]+nums[i] > 0 {
				tail--
			} else {
				head++
			}
		}
	}
	return resArr
}
