package Week_01

// https://leetcode-cn.com/problems/valid-parentheses/
// 2020-09-10

// 数组直接模拟
func isValid(s string) bool {
	stackArr := make([]byte, 0)
	parentMap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	// for range得到的不一样，为int32
	for i := 0; i < len(s); i++ {
		str := s[i]
		if _, ok := parentMap[str]; ok {
			stackArr = append(stackArr, str)
		} else {
			if len(stackArr) <= 0 || parentMap[stackArr[len(stackArr)-1]] != str {
				return false
			}
			// 出栈操作
			stackArr = stackArr[:len(stackArr)-1]
		}
	}
	if len(stackArr) > 0 {
		return false
	}
	return true
}

func isValid1(s string) bool {
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
