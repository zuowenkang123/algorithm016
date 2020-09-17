package Week_02

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

func preorder(root *Node) []int {
	res := []int{}

	if root == nil {
		return res
	}

	res = append(res, root.Val)
	for _, v := range root.Children {
		res = append(res, postorder(v)...)
	}

	return res
}

func levelOrder(root *Node) [][]int {

}
