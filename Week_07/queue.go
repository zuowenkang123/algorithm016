package Week_07

type TreeNodeQueue []*TreeNode

func (q *TreeNodeQueue) push(node *TreeNode) bool {
	*q = append(*q, node)
	return true
}

func (q *TreeNodeQueue) pop() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	first := []*TreeNode(*q)[0]
	*q = []*TreeNode(*q)[1:]
	return first
}

func (q *TreeNodeQueue) element() *TreeNode {
	first := []*TreeNode(*q)[0]
	return first
}

func (q *TreeNodeQueue) size() int {
	return len(*q)
}

func (q *TreeNodeQueue) pushAll(nodes []*TreeNode) bool {
	*q = append(*q, nodes...)
	return true
}

func (q *TreeNodeQueue) top() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	first := []*TreeNode(*q)[0]
	return first
}
