package Week_04

import "math"

// https://leetcode-cn.com/problems/word-ladder/
// 2020-09-28

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if indexOf(endWord, wordList) == -1 {
		return 0
	}
	step := 0
	visited := make([]bool, len(wordList))
	queue := []string{beginWord}
	for len(queue) > 0 {
		l := len(queue)
		step++
		for i := 0; i < l; i++ {
			if queue[i] == endWord {
				return step
			}
			for j := 0; j < len(wordList); j++ {
				if !visited[j] && oneDiff(queue[i], wordList[j]) {
					queue = append(queue, wordList[j])
					visited[j] = true
				}
			}
		}
		queue = queue[l:]
	}
	return 0
}

// 一步达到
func oneDiff(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}

func indexOf(str string, bank []string) int {
	for i, s := range bank {
		if str == s {
			return i
		}
	}
	return -1
}

var minChange1 int

// 同最小基因, 超时，由于wordList长度比较大
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	minChange1 = math.MaxInt32
	dfsLadderLength(beginWord, endWord, wordList, 1)
	if minChange1 == math.MaxInt32 {
		return 0
	}
	return minChange1
}

func dfsLadderLength(beginWord string, endWord string, wordList []string, change int) {
	if beginWord == endWord {
		minChange1 = min(minChange1, change)
		return
	}
	for i := 0; i < len(wordList); i++ {
		piece := wordList[i]
		if piece == "" {
			continue
		}
		diff := 0
		for j := 0; j < len(beginWord); j++ {
			if beginWord[j] != piece[j] {
				diff++
			}
		}
		if diff == 1 {
			wordList[i] = ""
			dfsLadderLength(piece, endWord, wordList, change+1)
			wordList[i] = piece
		}
	}
}
