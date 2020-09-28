package Week_04

import (
	"math"
)

// https://leetcode-cn.com/problems/minimum-genetic-mutation/
// 2020-09-28
var minChange int

func minMutation(start string, end string, bank []string) int {
	minChange = math.MaxInt32
	dfsMinMutation(start, end, bank, 0)
	if minChange == math.MaxInt32 {
		return -1
	}
	return minChange

}

func dfsMinMutation(start, end string, bank []string, change int) {
	if start == end {
		minChange = min(minChange, change)
		return
	}

	for i := 0; i < len(bank); i++ {
		piece := bank[i]
		// 已用过的片段
		if piece == "" {
			continue
		}
		// 获取基因库中不同为1的片段,作为改变一次后的新start
		diff := 0
		for j := 0; j < len(start); j++ {
			if start[j] != piece[j] {
				diff++
			}
		}
		if diff == 1 {
			// 置空,防止循环使用
			bank[i] = ""
			dfsMinMutation(piece, end, bank, change+1)
			// 还原回溯
			bank[i] = piece
		}
	}
}
