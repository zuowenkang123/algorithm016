package Week_02

// https://leetcode-cn.com/problems/remove-outermost-parentheses/submissions/
// 2020-09-16

// 计数法，通过count进行匹配，通过截取取掉外括号。
func removeOuterParentheses(S string) string {
	count := 0
	res := ""
	begin := 0
	for i := 0; i < len(S); i++ {
		if count == 0 {
			begin = i
		}
		if S[i] == '(' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res += S[begin+1 : i]
		}
	}
	return res
}

// 计数 另一种写法，对于这种舍弃第一个的写法，优先结果，最后操作。
func removeOuterParentheses1(S string) string {
	count := 0
	res := ""
	for i := 0; i < len(S); i++ {
		if S[i] == ')' {
			count--
		}
		if count > 0 {
			res += string(S[i])
		}
		if S[i] == '(' {
			count++
		}

	}
	return res
}

// 计数 可以更简单 实际就是左括号大于0就进，右括号大于1就进，然后根据括号进行count的标识处理
func removeOuterParentheses2(S string) string {
	count := 0
	res := ""
	for i := 0; i < len(S); i++ {
		if S[i] == '(' && count > 0 {
			res += string(S[i])
		}
		if S[i] == ')' && count > 1 {
			res += string(S[i])
		}
		if S[i] == '(' {
			count++
		} else {
			count--
		}
	}
	return res
}

// 计数 可以更简单
func removeOuterParentheses3(S string) string {
	count := 0
	res := ""
	for i := 0; i < len(S); i++ {
		if S[i] == '(' && count > 0 {
			res += string(S[i])
		}
		if S[i] == ')' && count > 1 {
			res += string(S[i])
		}
		if S[i] == '(' {
			count++
		} else {
			count--
		}
	}
	return res
}

// 栈
func removeOuterParentheses4(S string) string {
	stack := Stack{}
	res := ""
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			if false == stack.IsEmpty() {
				res += string(S[i])
			}
			stack = stack.Push(string(S[i]))
		}
		if S[i] == ')' {
			stack, _ = stack.Pop()
			if false == stack.IsEmpty() {
				res += string(S[i])
			}
		}
	}
	return res
}
