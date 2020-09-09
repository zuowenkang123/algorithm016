package Week_01

import "sort"

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// 四种题解：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/java-de-4chong-jie-fa-by-sdwwld/
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

// 临时数组
func findRepeatNumber2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 注意内存分配
	numArr := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		numArr[nums[i]]++
		if numArr[nums[i]] > 1 {
			return nums[i]
		}
	}
	return 0
}

// 下标归位法
func findRepeatNumber3(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return 0
}
