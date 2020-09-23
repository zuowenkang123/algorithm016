package Week_03

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// 2020-09-23

// 遍历导入法
func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	resRev := make([]int, 0, len(res))
	for i := len(res) - 1; i >= 0; i-- {
		resRev = append(resRev, res[i])
	}
	return resRev
}

// 栈
func reversePrint1(head *ListNode) []int {
	res := make([]int, 0)
	stack := StackInt{}
	for head != nil {
		stack.Push(head.Val)
		head = head.Next
	}

	resRev := make([]int, 0, len(res))
	for false == stack.IsEmpty() {
		resRev = append(resRev, stack.Pop())
	}
	return resRev
}

// 递归
var arr []int

func reversePrint2(head *ListNode) []int {
	arr = make([]int, 0)
	recur(head)
	res := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		res = append(res, arr[i])
	}
	return res
}

func recur(head *ListNode) {
	if head == nil {
		return
	}
	recur(head.Next)
	arr = append(arr, head.Val)

	return
}
