package Week_02

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
// 2020-09-16

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// 栈解法,装入root，不为空取出来，装入左边，装入右边
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := StackTreeNode([]*TreeNode{root}), []int{}
	for false == stack.IsEmpty() {
		cur := stack.Pop()
		res = append(res, cur.Val)

		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}
	return res
}
