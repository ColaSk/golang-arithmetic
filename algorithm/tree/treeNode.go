package tree

import "fmt"

// 二叉树 节点
// 为了简单将数据设为int型
type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (node *BinaryTreeNode) GetLeft() *BinaryTreeNode {
	return node.left
}

func (node *BinaryTreeNode) SetLeft(n *BinaryTreeNode) {
	node.left = n
}

func (node *BinaryTreeNode) GetRight() *BinaryTreeNode {
	return node.right
}

func (node *BinaryTreeNode) SetRight(n *BinaryTreeNode) {
	node.right = n
}

func (node *BinaryTreeNode) GetData() int {
	return node.data
}

func (node *BinaryTreeNode) SetData(data int) {
	node.data = data
}

func (node *BinaryTreeNode) Preorder() {

	if node == nil {
		return
	}
	fmt.Println(node.data, " ")
	node.left.Preorder()
	node.right.Preorder()
}

func NewBinNode(value int, left, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{
		data:  value,
		left:  left,
		right: right,
	}
}

// 为了简单 将 -1作为空
func CreateBinaryTree(node *BinaryTreeNode, len, loc int, values ...int) {

	if loc >= len {
		return
	}

	node.SetData(values[loc])

	if 2*loc+1 < len && values[2*loc+1] != -1 {

		node.SetLeft(&(BinaryTreeNode{}))
		CreateBinaryTree(node.GetLeft(), len, 2*loc+1, values...)
	}
	if 2*(loc+1) < len && values[2*(loc+1)] != -1 {

		node.SetRight(&(BinaryTreeNode{}))
		CreateBinaryTree(node.GetRight(), len, 2*(loc+1), values...)
	}

}
