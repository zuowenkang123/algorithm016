package Week_04

// https://leetcode-cn.com/problems/number-of-islands/
// 2020-09-28
var m int
var n int

func numIslands(grid [][]byte) int {
	count := 0
	n = len(grid)
	if n == 0 {
		return 0
	}
	m = len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfsMarking(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfsMarking(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'
	dfsMarking(grid, i+1, j)
	dfsMarking(grid, i-1, j)
	dfsMarking(grid, i, j+1)
	dfsMarking(grid, i, j-1)
}
