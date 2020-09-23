package Week_03

// https://leetcode-cn.com/problems/n-queens/
// 2020-09-23
import "strings"

var ans [][]string

var cols map[int]bool
var pie map[int]bool
var na map[int]bool

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{}
	}
	ans = make([][]string, 0, n)
	cols = make(map[int]bool, n)
	pie = make(map[int]bool, n)
	na = make(map[int]bool, n)
	dfsNQueens(n, 0, []int{})
	return ans
}

// 深度优先，先深入，然后回溯回来
func dfsNQueens(n int, row int, path []int) {
	if row >= n {
		ans = append(ans, generateBoard(path, n))
		return
	}
	for col := 0; col < n; col++ {
		// 终止条件
		if cols[col] || pie[col+row] || na[row-col] {
			continue
		}
		// 设置并下沉
		cols[col] = true
		pie[row+col] = true
		na[row-col] = true
		dfsNQueens(n, row+1, append(path, col))
		// 回溯
		cols[col] = false
		pie[row+col] = false
		na[row-col] = false
	}
	return
}

func generateBoard(path []int, n int) []string {
	var res []string
	for _, col := range path {
		res = append(res, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
	}
	return res
}
