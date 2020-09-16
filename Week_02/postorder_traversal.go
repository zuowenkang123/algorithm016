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
