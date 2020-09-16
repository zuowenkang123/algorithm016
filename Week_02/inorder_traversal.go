package Week_02

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// 2020-09-16

// 递归
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// 栈解法
func inorderTraversal1(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 取得栈顶元素
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// 栈解法, root不为空，左子树一致进栈，然后出栈，之后再放入
func inorderTraversal2(root *TreeNode) []int {
	stack := StackTreeNode{}
	res := []int{}
	for root != nil || false == stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		// 取得栈顶元素
		root = stack.Pop()
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
