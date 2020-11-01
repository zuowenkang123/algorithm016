package Week_07

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 2020-11-01

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := TreeNodeQueue{}
	queue.push(root)

	for queue.size() > 0 {
		size := queue.size()
		curLevel := make([]int, 0, size)
		for size > 0 {
			cur := queue.pop()
			if cur.Left != nil {
				queue.push(cur.Left)
			}
			if cur.Right != nil {
				queue.push(cur.Right)
			}
			curLevel = append(curLevel, cur.Val)
			size--
		}
		res = append(res, curLevel)
	}
	return res
}
