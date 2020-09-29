package Week_04

// https://leetcode-cn.com/problems/minesweeper/
// 2020-09-29
var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfsUpdateBoard(board, x, y)
	}
	return board
}
func dfsUpdateBoard(board [][]byte, x, y int) {
	cnt := 0
	// 遍历8个方位，查找周围的雷个数
	for i := 0; i < 8; i++ {
		tx, ty := x+dirX[i], y+dirY[i]
		if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
			continue
		}
		// 不用判断M，因为如果有 M 的话游戏已经结束了
		if board[tx][ty] == 'M' {
			cnt++
		}
	}

	if cnt > 0 {
		board[x][y] = byte(cnt + '0')
	} else {
		// 周围没有雷，需要递归查看周围
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			// 这里不需要在存在 B 的时候继续扩展，因为 B 之前被点击的时候已经被扩展过了
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
				continue
			}
			dfsUpdateBoard(board, tx, ty)
		}
	}
}
