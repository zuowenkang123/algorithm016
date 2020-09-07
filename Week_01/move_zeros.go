package Week_01

// https://leetcode-cn.com/problems/move-zeroes/

// 顺序让位法，同向指针
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i > pos { // #1
				nums[pos], nums[i] = nums[i], nums[pos]
			}
			pos++
		}
	}
	return
}

// 转存法, 新增了数组，不符合预期
func moveZeroes1(nums []int) {
	if len(nums) == 0 {
		return
	}
	numArr := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			numArr = append(numArr, nums[i])
		}
	}
	for j := 0; j <= len(nums)-len(numArr); j++ {
		numArr = append(numArr, 0)
	}

	nums = numArr
	return
}
