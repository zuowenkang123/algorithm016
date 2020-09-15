package Week_02

// https://leetcode-cn.com/problems/sliding-window-maximum/
// 2020-09-15

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	dequeue := list.New()
	resultArr := make([]int, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		// 如果超出队列则删除
		if dequeue.Len() > 0 && i-k >= dequeue.Front().Value.(int) {
			dequeue.Remove(dequeue.Front())
		}
		// 循环删除队列，直到遇到大的
		for dequeue.Len() > 0 && nums[i] > nums[dequeue.Back().Value.(int)] {
			dequeue.Remove(dequeue.Back())
		}
		// 每个数入栈
		dequeue.PushBack(i)

		// 存在满员则，获得最大加入数组
		if i >= k-1 {
			resultArr = append(resultArr, nums[dequeue.Front().Value.(int)])
		}
	}
	return resultArr
}
