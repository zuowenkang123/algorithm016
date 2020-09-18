package Week_02

import "container/heap"

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
// 2020-09-18

// 堆 求小数 大顶堆
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || k > len(arr) {
		return []int{}
	}
	h := &intMaxHeap{}
	heap.Init(h)
	for _, v := range arr {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := []int{}
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

// 堆, 取小数，小顶堆
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)
func getLeastNumbers1(arr []int, k int) []int {
	if k == 0 || k > len(arr) {
		return []int{}
	}
	h := &intMinHeap{}
	heap.Init(h)
	// 先全部加入
	for _, v := range arr {
		heap.Push(h, v)
	}

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

// 快排
func getLeastNumbers3(arr []int, k int) []int {
	quikSort(arr, 0, len(arr)-1)
	return arr[:k]
}

// 快排算法实现
func quikSort(arr []int, start, end int) {
	if start > end {
		return
	}
	mid := partition(arr, start, end)
	quikSort(arr, start, mid-1)
	quikSort(arr, mid+1, end)
}

func partition(arr []int, start, end int) int {
	pre := arr[end]
	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pre {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
