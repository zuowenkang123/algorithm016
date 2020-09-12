package Week_01

// https://leetcode-cn.com/problems/merge-two-sorted-lists/submissions/
// 2020-09-12

// 核心在于两个链表循环 O(m+n) O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	head.Next = l1
	pre := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next

	}

	if l2 != nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}
	return head.Next
}

// 递归
// 返回条件，就是l1或l2完结，返回小的list，同时递归后面的。O(m+n) O(m+n)
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists1(l2.Next, l1)
	return l2
}
