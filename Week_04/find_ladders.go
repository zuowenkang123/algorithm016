package Week_04

import "math"

// https://leetcode-cn.com/problems/word-ladder-ii/
// 2020-09-28

// 实在是困难，抄一个先
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ids := map[string]int{}
	for i, word := range wordList {
		ids[word] = i
	}
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList) - 1
	}
	if _, ok := ids[endWord]; !ok {
		return [][]string{}
	}
	n := len(wordList)
	edges := make([][]int, len(wordList))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	res := [][]string{}
	cost := make([]int, n)
	queue := [][]int{[]int{ids[beginWord]}}
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0
	for i := 0; i < len(queue); i++ {
		now := queue[i]
		last := now[len(now)-1]
		if last == ids[endWord] {
			temp := []string{}
			for _, index := range now {
				temp = append(temp, wordList[index])
			}
			res = append(res, temp)
		} else {
			for _, to := range edges[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last] + 1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return res
}

func transformCheck(form, to string) bool {
	for i := 0; i < len(form); i++ {
		if form[i] != to[i] {
			return form[i+1:] == to[i+1:]
		}
	}
	return false
}

// 错误写法
func findLadders1(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
	if indexOf(endWord, wordList) == -1 {
		return res
	}
	queue := []string{beginWord}
	visited := make([]bool, len(wordList))
	for len(queue) > 0 {
		l := len(queue)
		arr := make([]string, 0)
		arr = append(arr, queue[0])
		for i := 0; i < l; i++ {
			if queue[i] == endWord {
				res = append(res, arr)
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
	return res
}
