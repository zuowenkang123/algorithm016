package Week_02

import (
	"container/heap"
	"math"
)

// https://leetcode-cn.com/problems/ugly-number-ii/
// 2020-09-17

// 小顶堆
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

// dp
func nthUglyNumber1(n int) int {
	if n <= 0 {
		return -1
	}
	uglyNums := make([]int, n)
	uglyNums[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		m := min(uglyNums[p2]*2, uglyNums[p3]*3)
		m = min(m, uglyNums[p5]*5)
		uglyNums[i] = m
		if m == uglyNums[p2]*2 {
			p2++
		}
		if m == uglyNums[p3]*3 {
			p3++
		}
		if m == uglyNums[p5]*5 {
			p5++
		}
	}
	return uglyNums[n-1]
}
