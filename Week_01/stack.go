package Week_01

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
