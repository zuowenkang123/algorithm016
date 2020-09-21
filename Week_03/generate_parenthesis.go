package Week_03

// https://leetcode-cn.com/problems/generate-parentheses/
// 2020-09-21

var res []string

// 转变为n个盒子放括号问题
func generateParenthesis(n int) []string {
	res = make([]string, 0, 2*n)
	generate_parenthesis(0, 0, n, "")
	return res
}

// 首先按照2*n层进行划分，但是存在有效逻辑，细化为按照left和right划分。终止条件由level<2*n,变成left==n&&right==n
// 有效性细化为left处理和right处理，递归层次也变成交替处理
// 因此对于多模型，需要进行不同的处理
func generate_parenthesis(left, right, n int, str string) {
	// termination
	if left == n && right == n {
		res = append(res, str)
		return
	}

	// process
	// drill down
	if left < n {
		generate_parenthesis(left+1, right, n, str+"(")
	}
	if right < left {
		generate_parenthesis(left, right+1, n, str+")")
	}
	// reverse

	return

}

// 先判断后处理
func generate_parenthesis1(left, right, n int, str string) {
	if left < n || right > left {
		return
	}
	//if left == n && right == n {
	if left+right == 2*n {
		res = append(res, str)
		return
	}

	generate_parenthesis(left+1, right, n, str+"(")
	generate_parenthesis(left, right+1, n, str+")")

	return
}

// 减少法
func generateParenthesis1(n int) []string {
	res = make([]string, 0, 2*n)
	if n == 0 {
		return res
	}
	// 还剩多少个括号
	dfs(n, n, "")
	return res
}

// 先判断后处理
func dfs(left, right int, str string) {
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		res = append(res, str)
		return
	}
	if left > 0 {
		dfs(left-1, right, str+"(")
	}

	if right > 0 {
		dfs(left, right-1, str+")")
	}

	return
}
