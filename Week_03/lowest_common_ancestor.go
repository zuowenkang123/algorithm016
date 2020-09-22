package Week_03

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 2020-09-22

// 递归
// 最小在于自底向上，找的的第一个就是最小公共祖先
// 左右子树寻找p或q，找到上报父节点。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到则返回当前
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 左右子树存在
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

// path 搜索路径然后对比最小祖先
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}
