package Week_03

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{}
}

func rserialize(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = rserialize(root.Left, str)
		str = rserialize(root.Right, str)
	}
	return str
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return rserialize(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		if l[i] != "" {
			this.l = append(this.l, l[i])
		}
	}
	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}

	val, _ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val: val}
	this.l = this.l[1:]
	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()
	return root
}

// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/solution/shou-hui-tu-jie-gei-chu-dfshe-bfsliang-chong-jie-f/
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/solution/shen-du-you-xian-he-yan-du-you-xian-de-golangpytho/
