package Week_03

// https://leetcode-cn.com/problems/invert-binary-tree/submissions/
// 2020-09-21

// 递归
func invertTree(root *TreeNode) *TreeNode {
	// 终止
	if root == nil {
		return nil
	}
	// 处理加下沉
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

// 迭代
// 起始：装入root，开始非空循环
// 中间业务处理
// 非空入列继续循环
// 最后返回结果
func inverTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := TreeNodeQueue{}
	// 加入初始值，开始非空循环
	queue.push(root)
	for queue.size() != 0 {
		// 弹出处理
		node := queue.pop()
		// 交换
		node.Left, node.Right = node.Right, node.Left
		// 非空入列
		if node.Left != nil {
			queue.push(node.Left)
		}
		if node.Right != nil {
			queue.push(node.Right)
		}
	}
	return root
}
