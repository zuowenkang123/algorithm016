package Week_01

// https://leetcode-cn.com/problems/linked-list-cycle/
// 链表有环

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// map
func hasCycle1(head *ListNode) bool {
	listMap := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := listMap[head]; ok {
			return true
		}
		listMap[head] = true
		head = head.Next
	}
	return false
}
