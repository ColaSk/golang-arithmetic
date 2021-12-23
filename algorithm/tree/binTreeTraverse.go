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

func (node *BinaryTreeNode) Inorder() string {
	// 中序遍历

	var str string = ""

	if node == nil {
		return str
	}

	str += node.left.Inorder()
	str += strconv.Itoa(node.data)
	str += node.right.Inorder()

	return str
}

func (node *BinaryTreeNode) Postorder() string {
	var str string = ""

	if node == nil {
		return str
	}

	str += node.left.Postorder()
	str += node.right.Postorder()
	str += strconv.Itoa(node.data)

	return str
}
