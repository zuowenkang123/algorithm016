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
