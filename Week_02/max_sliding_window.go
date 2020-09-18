package Week_02

// https://leetcode-cn.com/problems/sliding-window-maximum/
// 2020-09-15

import (
	"container/list"
)

// 循环，判断是否超员，pk，加入，满则取头
func maxSlidingWindow(nums []int, k int) []int {
	dequeue := list.New()
	resultArr := make([]int, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		// 超员则删除老大
		if dequeue.Len() > 0 && i-k >= dequeue.Front().Value.(int) {
			dequeue.Remove(dequeue.Front())
		}
		// pk队列内，干掉小的
		for dequeue.Len() > 0 && nums[i] > nums[dequeue.Back().Value.(int)] {
			dequeue.Remove(dequeue.Back())
		}
		// 加入队列
		dequeue.PushBack(i)

		// 满员后，一直给出老大
		if i >= k-1 {
			resultArr = append(resultArr, nums[dequeue.Front().Value.(int)])
		}
	}
	return resultArr
}

// 暴力法，循环数组，内部比大小
func maxSlidingWindow1(nums []int, k int) []int {
	//边界条件判断
	if len(nums) == 0 {
		return []int{}
	}
	res := make([]int, len(nums)-k+1, len(nums)-k+1)
	for i := 0; i < len(nums)-k+1; i++ {
		maxInt := nums[i]
		//在每个窗口内找到最大值
		for j := 1; j < k; j++ {
			maxInt = max(maxInt, nums[i+j])
		}
		res[i] = maxInt
	}
	return res
}

// 双端扫描法
func maxSlidingWindow2(nums []int, k int) []int {
	return []int{}
}
