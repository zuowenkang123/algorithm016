package Week_01

import "sort"

// https://leetcode-cn.com/problems/two-sum/
// 2020-09-09
// 两数之和,问题在于如何获取原下标

// 排序+双指针+数组对比得下标
func twoSum(nums []int, target int) []int {
	// 条件判断
	if len(nums) < 2 {
		return []int{}
	}
	tmpArr := make([]int, len(nums))
	copy(tmpArr, nums)
	sort.Ints(tmpArr)
	head := 0
	tail := len(tmpArr) - 1
	numArr := make([]int, 0, 2)
	for head < tail {
		if tmpArr[head]+tmpArr[tail] == target {
			// 添加
			numArr = append(numArr, tmpArr[head], tmpArr[tail])
			// 处理左重复逻辑，右边的会被兼容
			for head < tail && tmpArr[head] == tmpArr[head+1] {
				head++
			}
			for head < tail && tmpArr[tail] == tmpArr[tail-1] {
				tail--
			}
			head++
			tail--
		} else if tmpArr[head]+tmpArr[tail] > target {
			tail--
		} else {
			head++
		}
	}
	indexArr := make([]int, 0)
	// 主数组在外层，保证每个元素仅一次
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(numArr); j++ {
			if nums[i] == numArr[j] {
				indexArr = append(indexArr, i)
				break
			}
		}
	}
	return indexArr
}

// 反向map法 O(n) O(n)
func twoSum1(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		// 只对应一个答案
		if _, ok := numMap[target-num]; ok {
			return []int{i, numMap[target-num]}
		}
		numMap[num] = i
	}
	return []int{}
}
