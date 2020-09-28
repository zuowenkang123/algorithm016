package Week_04

import "math"

// https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/
// 2020-09-28

// 采用广度优先，行处理模板
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := TreeNodeQueue{}
	queue.push(root)
	for queue.size() > 0 {
		size := queue.size()
		maxInt := math.MinInt32
		for size > 0 {
			node := queue.pop()
			size--
			maxInt = max(maxInt, node.Val)
			if node.Left != nil {
				queue.push(node.Left)
			}
			if node.Right != nil {
				queue.push(node.Right)
			}

		}
		res = append(res, maxInt)
	}
	return res
}
