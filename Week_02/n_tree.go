package Week_02

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
// N叉树的前序遍历
func preorder(root *Node) []int {
	res := []int{}

	if root == nil {
		return res
	}

	res = append(res, root.Val)
	// 充分体现了左右子树
	for _, v := range root.Children {
		res = append(res, postorder(v)...)
	}

	return res
}

// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
// 2020-09-17
// n叉树的后续遍历
func postorder(root *Node) []int {
	res := []int{}

	if root == nil {
		return res
	}

	for _, v := range root.Children {
		res = append(res, postorder(v)...)
	}

	res = append(res, root.Val)

	return res
}

var res [][]int

// N叉树的层级遍历 层级map动态规划解法
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *Node, level int) {
	if root == nil {
		return
	}
	if len(res) == level {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	for _, n := range root.Children {
		dfs(n, level+1)
	}
}

// N叉树的层级遍历 广度优先算法
func levelOrder1(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := NodeQueue{root}
	for queue.size() > 0 {
		levelArr := []int{}
		size := queue.size()
		for i := 0; i < size; i++ {
			node := queue.pop()
			levelArr = append(levelArr, node.Val)
			queue.pushAll(node.Children)
		}
		res = append(res, levelArr)
	}
	return res
}
