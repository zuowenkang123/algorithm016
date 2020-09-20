package Week_02

// 字符串栈
type Stack []string

func (s Stack) Push(v string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, string) {
	l := len(s)
	if l == 0 {
		return s, ""
	}
	return s[:l-1], s[l-1]
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// int栈
type StackInt []int

func (s *StackInt) Push(v int) {
	*s = append(*s, v)
}

func (s *StackInt) Top() int {
	l := len(*s)
	if l == 0 {
		return 0
	}
	return (*s)[l-1]
}

func (s *StackInt) Pop() int {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

func (s *StackInt) IsEmpty() bool {
	return len(*s) == 0
}

// 树栈
type StackTreeNode []*TreeNode

func (s *StackTreeNode) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *StackTreeNode) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

func (s *StackTreeNode) Top() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	return n
}

func (s StackTreeNode) IsEmpty() bool {
	return len(s) == 0
}
