package Week_03

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 2020-09-22
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/xiong-mao-shua-ti-python3-xian-xu-zhao-gen-hua-fen/

// 生成树
// inorder.index(preorder[0]) 这一步获取根的索引值，题目说树中的各个节点的值都不相同，也确保了这步得到的结果是唯一准确的。而且这个idx还能当长度用相当于 左+根 的长度，因为 左+根 和 根+左 是等长的。
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 取根节点
	root := &TreeNode{preorder[0], nil, nil}
	// 得到root的index
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	// 然后左边为左边，右边为右边
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

var inMap map[int]int

func buildTree2(preorder []int, inorder []int) *TreeNode {
	for i := 0; i < len(inorder); i++ {
		inMap[inorder[i]] = i
	}
	return dp(preorder, 0, len(inorder)-1)
}

var index int

func dp(preorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	i := inMap[preorder[index]]
	root := &TreeNode{Val: preorder[index]}
	index++
	root.Left = dp(preorder, left, i-1)
	root.Right = dp(preorder, i+1, right)
	return root
}
