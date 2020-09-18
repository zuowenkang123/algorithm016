package Week_02

type intMaxHeap []int

func (h intMaxHeap) Len() int {
	return len(h)
}

func (h intMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intMaxHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *intMaxHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

type intMinHeap []int

func (h intMinHeap) Len() int {
	return len(h)
}

func (h intMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intMinHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *intMinHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
