package Week_03

// https://leetcode-cn.com/problems/merge-two-binary-trees/
// 2020-09-23

// 递归，当前即是所有。当前判断t1、t2为空，即所有逻辑都是，当前赋值左右节点，即所有都是。
// dfs
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

// bfs
func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	queue := TreeNodeQueue{}
	queue.push(t1)
	queue.push(t2)
	for queue.size() > 0 {
		node1 := queue.pop()
		node2 := queue.pop()
		// 计算逻辑
		node1.Val += node2.Val
		if node1.Left == nil {
			// 追加逻辑
			node1.Left = node2.Left
		} else {
			if node2.Left != nil {
				queue.push(node1.Left)
				queue.push(node2.Left)
			}
		}
		if node1.Right == nil {
			// 追加逻辑
			node1.Right = node2.Right
		} else {
			if node2.Right != nil {
				queue.push(node1.Right)
				queue.push(node2.Right)
			}
		}
	}
	return t1
}
