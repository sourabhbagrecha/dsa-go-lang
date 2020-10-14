package binarysearchtree

import (
	"fmt"
)

type leaf struct {
	data  int
	left  *leaf
	right *leaf
}

//BinarySearchTree Data Structure
type BinarySearchTree struct {
	root *leaf
}

//Insert a new leaf to the BST
func (b *BinarySearchTree) Insert(data int) *BinarySearchTree {
	newLeaf := &leaf{data: data}
	if b.root == nil {
		b.root = newLeaf
		return b
	}
	currentLeaf := b.root
	for true {
		if data < currentLeaf.data {
			if currentLeaf.left == nil {
				currentLeaf.left = newLeaf
				return b
			}
			currentLeaf = currentLeaf.left
		} else if data > currentLeaf.data {
			if currentLeaf.right == nil {
				currentLeaf.right = newLeaf
				return b
			}
			currentLeaf = currentLeaf.right
		} else {
			return b
		}
	}
	return b
}

//BFSTraversal traverse all the nodes in the BST in a Breadth First Search order
func (b *BinarySearchTree) BFSTraversal() []int {
	visited := []int{}

	if b.root == nil {
		return visited
	}

	consideredQueue := []leaf{}
	node := *b.root
	consideredQueue = append(consideredQueue, node)
	for len(consideredQueue) > 0 {
		node = consideredQueue[0]
		consideredQueue = consideredQueue[1:]
		visited = append(visited, node.data)
		if node.left != nil {
			consideredQueue = append(consideredQueue, *node.left)
		}
		if node.right != nil {
			consideredQueue = append(consideredQueue, *node.right)
		}
	}
	return visited
}

func preOrderHelper(node *leaf, preOrder []int) []int {
	//visit root
	preOrder = append(preOrder, node.data)
	//traverse left
	if node.left != nil {
		preOrder = preOrderHelper(node.left, preOrder)
	}
	//traverse right
	if node.right != nil {
		preOrder = preOrderHelper(node.right, preOrder)
	}
	return preOrder
}

func inOrderHelper(node *leaf, inOrder []int) []int {
	//traverse left
	if node.left != nil {
		inOrder = inOrderHelper(node.left, inOrder)
	}
	//visit root
	inOrder = append(inOrder, node.data)
	//traverse right
	if node.right != nil {
		inOrder = inOrderHelper(node.right, inOrder)
	}
	return inOrder
}

func postOrderHelper(node *leaf, postOrder []int) []int {
	//traverse left
	if node.left != nil {
		postOrder = postOrderHelper(node.left, postOrder)
	}
	//traverse right
	if node.right != nil {
		postOrder = postOrderHelper(node.right, postOrder)
	}
	//visit root
	postOrder = append(postOrder, node.data)
	return postOrder
}

//DFSTraversal traverse all the nodes in the BST in the Depth First Search(Pre-order) order
func (b *BinarySearchTree) DFSTraversal(orderType string) []int {
	var order []int
	switch orderType {
	case "pre-order":
		order = preOrderHelper(b.root, order)
		break
	case "in-order":
		order = inOrderHelper(b.root, order)
		break
	case "post-order":
		order = postOrderHelper(b.root, order)
		break
	default:
		order = preOrderHelper(b.root, order)
		break
	}
	return order
}

//Print the BST
func (b *BinarySearchTree) Print() {
	currentLeaf := b.root
	printBSTFromANode(currentLeaf)
}

func printBSTFromANode(l *leaf) *leaf {
	var left, right *leaf
	if l.left != nil {
		left = printBSTFromANode(l.left)
	}
	if l.right != nil {
		right = printBSTFromANode(l.right)
	}
	if l.left != nil || l.right != nil {
		fmt.Println("Traversing", left, right, l)
	}
	return l
}

//Find a leaf with the given data in BST
func (b *BinarySearchTree) Find(data int) bool {
	if b.root == nil {
		return false
	}
	currentLeaf := b.root
	for true {
		if data == currentLeaf.data {
			return true
		} else if data < currentLeaf.data {
			if currentLeaf.left == nil {
				return false
			}
			currentLeaf = currentLeaf.left
		} else {
			if currentLeaf.right == nil {
				return false
			}
			currentLeaf = currentLeaf.right
		}
	}
	return false
}
