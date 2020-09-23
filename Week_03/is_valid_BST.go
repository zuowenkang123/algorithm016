package Week_03

import "math"

// https://leetcode-cn.com/problems/validate-binary-search-tree/
// 2020-09-21

// 验证二叉树
// 需要传递判断,对于左子树要求小于等于当前值，对于右子树要求大于等于当前值，同时需要最大最小边界
func isValidBST(root *TreeNode) bool {
	return is_valid_BST(root, math.MinInt64, math.MaxInt64)

}

// 递归，当前的值如果不符合，则所有都不符合，同时判断左右子树。
func is_valid_BST(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return is_valid_BST(root.Left, lower, root.Val) && is_valid_BST(root.Right, root.Val, upper)
}

// 中序遍历，二叉搜索树
var lower int

func isValidBST1(root *TreeNode) bool {
	lower = math.MinInt64
	return is_Valid_BST1(root)

}

// 初始lower，然后先左子树对比，然后当前对比，然后右子树对比
func is_Valid_BST1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 左子树不符合则直接返回
	if false == is_Valid_BST1(root.Left) {
		return false
	}
	// 当前值不负责，返回
	if root.Val <= lower {
		return false
	}
	// 更新lower
	lower = root.Val
	// 遍历右子树
	return is_Valid_BST1(root.Right)
}

// 迭代,一左到底，处理中，最后累算右，循环
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	lower := math.MinInt64
	for len(stack) > 0 || root != nil {
		// 入栈左子树数据
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹出当前
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 比较数据
		if root.Val <= lower {
			return false
		}
		// 写入数据
		lower = root.Val

		// 更换右子树
		root = root.Right
	}
	return true
}
