package Week_02

// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/submissions/
// 2020-09-16

// 递归
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 栈
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	slice := make([]int, 0)
	for len(stack) > 0 {
		root := stack[len(stack)-1]
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Left == nil && root.Right == nil {
			stack = stack[:len(stack)-1]
			slice = append(slice, root.Val)
		}
		root.Left = nil
		root.Right = nil
	}
	return slice
}

// 栈，取头，右子树进栈，左子树进栈，为空当前取出，设置左右为空
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := StackTreeNode([]*TreeNode{root}), []int{}
	for false == stack.IsEmpty() {
		root := stack.Top()
		if root.Right != nil {
			stack.Push(root.Right)
		}
		if root.Left != nil {
			stack.Push(root.Left)
		}
		if root.Left == nil && root.Right == nil {
			stack.Pop()
			res = append(res, root.Val)
		}
		root.Left = nil
		root.Right = nil
	}
	return res
}
