package Week_01

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 2020-09-10

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	// 必须要有一个前置节点，否则交换的时候无法找到前置节点
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next
		// 前面连接
		pre.Next = second
		// 简化为second.Next, first.Next = first, second.Next
		tmp := second.Next
		second.Next = first
		first.Next = tmp
		// 移动
		pre = first // 这个地方得用first作为前置，因为经过上面交换这个变成了下一组的前置节点 移动
	}
	return dummy.Next
}

// 递归
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 将当前给next，并最终返回next
	next := head.Next
	// head连接下一组
	head.Next = swapPairs(next.Next)
	// next指向head
	next.Next = head
	return next
}
