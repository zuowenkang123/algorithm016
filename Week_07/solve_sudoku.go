package Week_07

// https://leetcode-cn.com/problems/sudoku-solver/
// 2020-11-01
func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}
	solve(&board)
}

func solve(board *[][]byte) bool {
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[0]); j++ {
			if (*board)[i][j] == '.' {
				for c := '1'; c <= '9'; c++ {
					if isValid((*board), i, j, byte(c)) {
						(*board)[i][j] = byte(c)
						if solve(board) {
							return true
						} else {
							(*board)[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		box_row := (row/3)*3 + i/3
		bow_col := (col/3)*3 + i%3
		if board[box_row][bow_col] != '.' && board[box_row][bow_col] == c {
			return false
		}
	}
	return true
}
