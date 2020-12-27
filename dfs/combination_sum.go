package dfs

// https://leetcode-cn.com/problems/combination-sum/submissions/

var resComSum [][]int

func combinationSum(candidates []int, target int) [][]int {
	resComSum = make([][]int, 0)
	CombinationSumDfs(0, []int{}, 0, target, candidates)
	return resComSum
}

func CombinationSumDfs(start int, combine []int, sum, target int, candidates []int) {
	if sum > target {
		return
	}
	if sum == target {
		tmp := make([]int, len(combine))
		copy(tmp, combine)
		resComSum = append(resComSum, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		combine = append(combine, candidates[i])
		CombinationSumDfs(i, combine, sum+candidates[i], target, candidates)
		combine = combine[:len(combine)-1]
	}
}
