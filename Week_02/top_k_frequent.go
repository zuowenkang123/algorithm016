package Week_02

import (
	"container/heap"
	"math/rand"
)

// https://leetcode-cn.com/problems/top-k-frequent-elements/
// 2020-09-18

// 堆方式
func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	// 小顶堆
	for key, value := range numMap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, h.Len())
	// 正序排列
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// 快排逻辑
func topKFrequent1(nums []int, k int) []int {
	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num]++
	}
	pairs := [][]int{}
	for key, value := range numMap {
		pairs = append(pairs, []int{key, value})
	}
	left := 0
	right := len(pairs) - 1
	index := k - 1
	for left <= right {
		pos := partitionPair(pairs, left, right)
		if pos == index {
			break
		} else if index < pos {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}
	res := make([]int, 0)
	for i := len(pairs) - 1; i >= len(pairs)-k; i-- {
		res = append(res, pairs[i][0])
	}
	return res
}

func partitionPair(pair [][]int, left, right int) int {
	randInt := rand.Int() % (right - left + 1)
	pair[left], pair[randInt] = pair[randInt], pair[left]
	pos := left
	posValue := pair[left][1]

	for i := left + 1; i <= right; i++ {
		if pair[i][1] < posValue {
			pos++
			pair[i], pair[pos] = pair[pos], pair[i]
		} else {
			pair[left], pair[pos] = pair[pos], pair[left]
		}
	}

	return pos
}
