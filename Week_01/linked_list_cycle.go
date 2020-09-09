package Week_01

// https://leetcode-cn.com/problems/linked-list-cycle/
// 链表有环

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next
	for fast != slow {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
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
