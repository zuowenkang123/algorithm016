package Week_02

import "container/heap"

// 堆
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || k > len(arr) {
		return []int{}
	}
	h := &intHeap{}
	heap.Init(h)
	for _, v := range arr {
		if h.Len() < k {
			heap.Push(h, v)
		} else {
			if (*h)[0] > v {
				heap.Pop(h)
				heap.Push(h, v)
			}
		}
	}

	res := []int{}
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

// 排序，取前k个
func getLeastNumbers1(arr []int, k int) []int {
	return []int{}
}

// 快排
func getLeastNumbers2(arr []int, k int) []int {
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
