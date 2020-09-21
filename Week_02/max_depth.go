package Week_02

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/submissions/
// 2020-09-20

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 与广度优先不同，需要在装入下一层是，清空上一层，以此累加
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		// pop出当前层数据
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

// 与广度优先不同，需要在装入下一层是，清空上一层，以此累加
func maxDepth2(root *TreeNode) int {
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
