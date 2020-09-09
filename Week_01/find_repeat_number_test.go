package Week_01

import "sort"

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// 2020-09-09

// map是常用的去重
func findRepeatNumber(nums []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return num
		}
		numMap[num] = true
	}
	return 0
}

// 有序的去重,后面对比前面
func findRepeatNumber1(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}
