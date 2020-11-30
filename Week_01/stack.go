package Week_01

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	l := len(*s)
	if l == 0 {
		return ""
	}
	str := (*s)[l-1]
	*s = (*s)[:l-1]
	return str
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

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
	l := len(*s)
	if l == 0 {
		return 0
	}
	n := (*s)[l-1]
	*s = (*s)[:l-1]
	return n
}

func (s *StackInt) IsEmpty() bool {
	return len(*s) == 0
}
