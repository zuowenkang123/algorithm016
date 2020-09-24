package Week_03

// https://leetcode-cn.com/problems/permutations/
// 2020-09-23

var ansPermute [][]int

func permute(nums []int) [][]int {
	ansPermute = make([][]int, 0)
	backtrack(len(nums), 0, nums)
	return ansPermute

}

// 基于交换方法
func backtrack(n, first int, nums []int) {
	if first == n {
		arr := make([]int, len(nums))
		copy(arr, nums)
		ansPermute = append(ansPermute, arr)
		return
	}
	for i := first; i < n; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtrack(n, first+1, nums)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

var used []bool // 需要时数组，map存在不存在情况

func permute1(nums []int) [][]int {
	ansPermute = make([][]int, 0)
	path := make([]int, 0)
	used = make([]bool, len(nums))
	backtrack1(nums, path)
	return ansPermute

}

// 基于map方法
func backtrack1(nums []int, path []int) {
	if len(nums) == len(path) {
		arr := make([]int, len(path))
		copy(arr, path)
		ansPermute = append(ansPermute, arr)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrack1(nums, path)
		used[i] = false
		path = path[:len(path)-1]
	}
}
