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
			// 数组小于0，或者数组不小于0，栈首数据不相等
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
	for _, val := range s {
		str := string(rune(val))
		if str == "(" || str == "[" || str == "{" {
			stack.Push(str)
		} else {
			if _, ok := parentMap[str]; !ok {
				return false
			}
			top := stack.Pop()
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

var parentMap = map[int32]int32{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid2(s string) bool {
	stackArr := make([]int32, 0)
	for _, v := range s {
		// 左括号入栈
		if _, ok := parentMap[v]; ok {
			stackArr = append(stackArr, v)
		} else {
			// 不符合返回
			if len(stackArr) == 0 || v != parentMap[stackArr[len(stackArr)-1]] {
				return false
			}
			// 匹配到出栈
			stackArr = stackArr[:len(stackArr)-1]
		}
	}
	// 剩余数据
	if len(stackArr) > 0 {
		return false
	}
	return true
}
