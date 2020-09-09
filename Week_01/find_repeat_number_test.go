package Week_01

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
