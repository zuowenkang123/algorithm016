package Week_03

// https://leetcode-cn.com/problems/subsets/
// 2020-09-24
var resSubsets [][]int
var tmpSubsets []int

func subsets(nums []int) [][]int {
	resSubsets = make([][]int, 0)
	tmpSubsets = []int{}
	dfsSubsets(nums, 0)
	return resSubsets
}

// 不需要判断添加条件
func dfsSubsets(nums []int, start int) {
	arr := make([]int, len(tmpSubsets))
	copy(arr, tmpSubsets)
	resSubsets = append(resSubsets, arr)

	for i := start; i < len(nums); i++ {
		tmpSubsets = append(tmpSubsets, nums[i])
		dfsSubsets(nums, i+1)
		tmpSubsets = tmpSubsets[:len(tmpSubsets)-1]
	}
}
