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
