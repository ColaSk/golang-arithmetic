package tree

import "strconv"

func (node *BinaryTreeNode) Preorder() string {
	// 先序遍历
	var str string = ""

	if node == nil {
		return str
	}

	str += strconv.Itoa(node.data)
	str += node.left.Preorder()
	str += node.right.Preorder()

	return str
}
