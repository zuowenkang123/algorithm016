package Week_03

// https://leetcode-cn.com/problems/merge-two-binary-trees/
// 2020-09-23

// 递归，当前即是所有。当前判断t1、t2为空，即所有逻辑都是，当前赋值左右节点，即所有都是。
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1

}

// 错误示例
func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	root := &TreeNode{}
	for t1 != nil || t2 != nil {
		if t1 != nil && t2 != nil {
			root.Val = t1.Val + t2.Val
		} else if t1 != nil {
			root.Val = t1.Val
		} else if t2 != nil {
			root.Val = t2.Val
		}
		mergeTrees(t1.Left, t2.Left)
		mergeTrees(t1.Right, t2.Right)
	}
	return root
}
