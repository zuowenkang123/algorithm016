package Week_02

import (
	"container/heap"
	"math"
)

// https://leetcode-cn.com/problems/ugly-number-ii/
// 2020-09-17

// 最小堆
func nthUglyNumber(n int) int {
	uglyNum := 0
	h := intMinHeap{}
	heap.Init(&h)
	heap.Push(&h, 1)
	for n > 0 {
		for h[0] == uglyNum {
			heap.Pop(&h)
		}
		uglyNum = heap.Pop(&h).(int)
		// 将后面的丑数加入堆
		for _, num := range []int{2, 3, 5} {
			if num*uglyNum <= math.MaxInt32 {
				heap.Push(&h, num*uglyNum)
			}
		}
		n--
	}
	return uglyNum
}
