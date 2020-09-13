package Week_01

type MyCircularDeque struct {
	nums     []int
	capacity int
	lenth    int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		nums:     make([]int, 0, 10),
		capacity: k,
		lenth:    0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.lenth < this.capacity {
		this.nums = append([]int{value}, this.nums...)
		this.lenth++
		return true
	}
	return false
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.lenth < this.capacity {
		this.nums = append(this.nums, value)
		this.lenth++
		return true
	}
	return false
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.lenth > 0 {
		this.nums = this.nums[1:]
		this.lenth--
		return true
	}
	return false
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.lenth > 0 {
		this.nums = this.nums[:this.lenth-1]
		this.lenth--
		return true
	}
	return false
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.lenth > 0 {
		return this.nums[0]
	}
	return -1
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.lenth > 0 {
		return this.nums[this.lenth-1]
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.lenth == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.lenth == this.capacity {
		return true
	}
	return false
}
