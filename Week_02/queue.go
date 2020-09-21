package Week_02

type NodeQueue []*Node

func (q *NodeQueue) push(node *Node) bool {
	*q = append(*q, node)
	return true
}

func (q *NodeQueue) pop() *Node {
	if len(*q) == 0 {
		return nil
	}
	first := []*Node(*q)[0]
	*q = []*Node(*q)[1:]
	return first
}

func (q *NodeQueue) element() *Node {
	first := []*Node(*q)[0]
	return first
}

func (q *NodeQueue) size() int {
	return len(*q)
}

func (q *NodeQueue) pushAll(nodes []*Node) bool {
	*q = append(*q, nodes...)
	return true
}

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
