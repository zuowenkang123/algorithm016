package Week_01

import "sort"

// https://leetcode-cn.com/problems/two-sum/
// 2020-09-09
// 两数之和,问题在于如何获取原下标【map不行，相同的数无法区分】
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
	pos := 0
	indexArr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if pos >= 2 {
			break
		}
		if numArr[pos] == nums[i] {
			indexArr = append(indexArr, i)
			pos++
		}
	}
	return indexArr
}
