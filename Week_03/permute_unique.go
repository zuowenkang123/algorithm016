package Week_03

import "sort"

// https://leetcode-cn.com/problems/permutations-ii/
// 2020-09-24
var ansPermuteUnique [][]int
var usedUnique []bool

func permuteUnique(nums []int) [][]int {
	ansPermuteUnique = make([][]int, 0)
	path := make([]int, 0)
	usedUnique = make([]bool, len(nums))
	sort.Ints(nums)
	backtrack_unique(nums, path)
	return ansPermuteUnique
}

// 基于map方法
func backtrack_unique(nums []int, path []int) {
	if len(nums) == len(path) {
		arr := make([]int, len(path))
		copy(arr, path)
		ansPermuteUnique = append(ansPermuteUnique, arr)
		return
	}
	for i := 0; i < len(nums); i++ {
		if usedUnique[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !usedUnique[i-1] {
			continue
		}
		usedUnique[i] = true
		path = append(path, nums[i])
		backtrack_unique(nums, path)
		usedUnique[i] = false
		path = path[:len(path)-1]

	}
}
