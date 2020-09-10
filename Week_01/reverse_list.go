package Week_01

// https://leetcode-cn.com/problems/reverse-linked-list/
// 2020-09-10

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		// 	curr.Next, pre, curr = pre, curr, curr.Next
		tmp := curr.Next
		// 核心做了反向
		curr.Next = pre
		// 后面两个做了移动
		pre = curr
		curr = tmp
	}
	return pre
}

// 递归写法
func reverse_list(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := reverse_list(head.Next)

	head.Next.Next = head
	head.Next = nil
	return curr
}
