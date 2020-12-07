package Week_03

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
// 2020-09-21
// 递归 最小高度为当前+左右子树最小高度
// 4种情况
func minDepth(root *TreeNode) int {
	if root == nil { // 都为空
		return 0
	}
	dl := minDepth(root.Left)
	dr := minDepth(root.Right)

	if root.Right != nil && root.Left == nil { // 右子树不为空
		return dr + 1
	} else if root.Left != nil && root.Right == nil { // 左子树不为空
		return dl + 1
	} else { // 都不为空
		return min(dl, dr) + 1
	}
}

// 迭代，广度搜索，需要删除当前层处理
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := TreeNodeQueue{}
	queue.push(root)
	depth := 0
	for queue.size() > 0 {
		depth++
		size := queue.size()
		// pop出当前层数据
		for size > 0 {
			node := queue.pop()
			size--
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue.push(node.Left)
			}
			if node.Right != nil {
				queue.push(node.Right)
			}
		}
	}
	return depth
}

// 迭代，广度搜索，需要删除当前层处理
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		size := len(queue)
		// pop出当前层数据
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			size--
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}
