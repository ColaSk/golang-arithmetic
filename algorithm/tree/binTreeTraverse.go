package tree

import (
	"container/list"
	"strconv"
)

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
	// 后序遍历
	var str string = ""

	if node == nil {
		return str
	}

	str += node.left.Postorder()
	str += node.right.Postorder()
	str += strconv.Itoa(node.data)

	return str
}

func (node *BinaryTreeNode) BreadthFirst() string {
	// 广度优先遍历

	queue := list.New()

	str := ""

	if node == nil {
		return str
	}

	queue.PushBack(node)

	for queue.Len() != 0 {
		temp := queue.Front()
		queue.Remove(temp)

		// go 断言
		if tempNode, ok := temp.Value.(*BinaryTreeNode); ok {

			str += strconv.Itoa(tempNode.data)

			if tempNode.left != nil {
				queue.PushBack(tempNode.left)
			}

			if tempNode.right != nil {
				queue.PushBack(tempNode.right)
			}
		}

	}

	return str

}
