package Week_07

// https://leetcode-cn.com/problems/valid-sudoku/description/
// 2020-11-01 20:37:19

func isValidSudoku(board [][]byte) bool {
	var row, col, sbox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				// 行
				if row[i][num] == 1 {
					return false
				} else {
					row[i][num] = 1
				}
				// 列
				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
				// 块
				box_index := (i/3)*3 + j/3
				if sbox[box_index][num] == 1 {
					return false
				} else {
					sbox[box_index][num] = 1
				}

			}
		}
	}
	return true
}
