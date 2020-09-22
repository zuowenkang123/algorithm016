package Week_03

// https://leetcode-cn.com/problems/combinations/
// 2020-09-22

// 递归
// 选第n个元素：装入n-1,k-1，再加入n
// 不选第n个元素：装入n-1,k
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if n < k || k == 0 {
		return res
	}
	res = combine(n-1, k-1)
	if len(res) == 0 {
		arr := make([]int, 0)
		res = append(res, arr)
	}
	for i, _ := range res {
		res[i] = append(res[i], n)
	}

	res = append(res, combine(n-1, k)...)
	return res
}

// 回溯
var resInt [][]int

func combine1(n int, k int) [][]int {
	resInt = make([][]int, 0)
	dfs1(n, k, 1, []int{})
	return resInt
}

// 从start开始，直到k用完
func dfs1(n, k, start int, arr []int) {
	if k == 0 {
		// 需要copy
		comb := make([]int, len(arr))
		copy(comb, arr)
		resInt = append(resInt, comb)
		return
	}
	for i := start; i <= n-k+1; i++ {
		arr = append(arr, i)
		// 装入当前后，下一层k-1,i+1
		dfs1(n, k-1, i+1, arr)
		// 尝试装完后，回溯处理上一层
		arr = arr[:len(arr)-1]
	}
}
