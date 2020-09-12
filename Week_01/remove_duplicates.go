package Week_01

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 2020-09-12
// 双指针法 O(n) O(1)
func removeDuplicates(nums []int) int {
	pos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			// 交换下一个
			pos++
			nums[pos] = nums[i]
		}
	}
	return pos + 1 // 给个数+1
}

// 优化，同位置不交换
func removeDuplicates1(nums []int) int {
	pos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			// 交换下一个
			pos++
			if pos != i { // 非同一个位置才交换
				nums[pos] = nums[i]
			}
		}
	}
	return pos + 1 // 给个数+1
}
