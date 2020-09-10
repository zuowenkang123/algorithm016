package Week_01

// https://leetcode-cn.com/problems/valid-parentheses/
// 2020-09-10
func isValid(s string) bool {
	stack := Stack{}
	parentMap := make(map[string]string)
	parentMap[")"] = "("
	parentMap["]"] = "["
	parentMap["}"] = "{"
	var top string
	for _, val := range s {
		str := string(rune(val))
		if str == "(" || str == "[" || str == "{" {
			stack = stack.Push(str)
		} else {
			if _, ok := parentMap[str]; !ok {
				return false
			}
			stack, top = stack.Pop()
			if top != parentMap[str] {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

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
