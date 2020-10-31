package binarysearchtree

import (
	"fmt"
)

// Leaf Data structure
type Leaf struct {
	Data  int
	Left  *Leaf
	Right *Leaf
}

//BinarySearchTree Data Structure
type BinarySearchTree struct {
	Root *Leaf
}

//Insert a new Leaf to the BST
func (b *BinarySearchTree) Insert(Data int) *BinarySearchTree {
	newLeaf := &Leaf{Data: Data}
	if b.Root == nil {
		b.Root = newLeaf
		return b
	}
	currentLeaf := b.Root
	for true {
		if Data < currentLeaf.Data {
			if currentLeaf.Left == nil {
				currentLeaf.Left = newLeaf
				return b
			}
			currentLeaf = currentLeaf.Left
		} else if Data > currentLeaf.Data {
			if currentLeaf.Right == nil {
				currentLeaf.Right = newLeaf
				return b
			}
			currentLeaf = currentLeaf.Right
		} else {
			return b
		}
	}
	return b
}

//BFSTraversal traverse all the nodes in the BST in a Breadth First Search order
func (b *BinarySearchTree) BFSTraversal() []int {
	visited := []int{}

	if b.Root == nil {
		return visited
	}

	consideredQueue := []Leaf{}
	node := *b.Root
	consideredQueue = append(consideredQueue, node)
	for len(consideredQueue) > 0 {
		node = consideredQueue[0]
		consideredQueue = consideredQueue[1:]
		visited = append(visited, node.Data)
		if node.Left != nil {
			consideredQueue = append(consideredQueue, *node.Left)
		}
		if node.Right != nil {
			consideredQueue = append(consideredQueue, *node.Right)
		}
	}
	return visited
}

func preOrderHelper(node *Leaf, preOrder []int) []int {
	//visit Root
	preOrder = append(preOrder, node.Data)
	//traverse Left
	if node.Left != nil {
		preOrder = preOrderHelper(node.Left, preOrder)
	}
	//traverse Right
	if node.Right != nil {
		preOrder = preOrderHelper(node.Right, preOrder)
	}
	return preOrder
}

func inOrderHelper(node *Leaf, inOrder []int) []int {
	//traverse Left
	if node.Left != nil {
		inOrder = inOrderHelper(node.Left, inOrder)
	}
	//visit Root
	inOrder = append(inOrder, node.Data)
	//traverse Right
	if node.Right != nil {
		inOrder = inOrderHelper(node.Right, inOrder)
	}
	return inOrder
}

func postOrderHelper(node *Leaf, postOrder []int) []int {
	//traverse Left
	if node.Left != nil {
		postOrder = postOrderHelper(node.Left, postOrder)
	}
	//traverse Right
	if node.Right != nil {
		postOrder = postOrderHelper(node.Right, postOrder)
	}
	//visit Root
	postOrder = append(postOrder, node.Data)
	return postOrder
}

//DFSTraversal traverse all the nodes in the BST in the Depth First Search(Pre-order) order
func (b *BinarySearchTree) DFSTraversal(orderType string) []int {
	var order []int
	switch orderType {
	case "pre-order":
		order = preOrderHelper(b.Root, order)
		break
	case "in-order":
		order = inOrderHelper(b.Root, order)
		break
	case "post-order":
		order = postOrderHelper(b.Root, order)
		break
	default:
		order = preOrderHelper(b.Root, order)
		break
	}
	return order
}

//Print the BST
func (b *BinarySearchTree) Print() {
	currentLeaf := b.Root
	printBSTFromANode(currentLeaf)
}

func printBSTFromANode(l *Leaf) *Leaf {
	var Left, Right *Leaf
	if l.Left != nil {
		Left = printBSTFromANode(l.Left)
	}
	if l.Right != nil {
		Right = printBSTFromANode(l.Right)
	}
	if l.Left != nil || l.Right != nil {
		fmt.Println("Traversing", Left, Right, l)
	}
	return l
}

//Find a Leaf with the given Data in BST
func (b *BinarySearchTree) Find(Data int) bool {
	if b.Root == nil {
		return false
	}
	currentLeaf := b.Root
	for true {
		if Data == currentLeaf.Data {
			return true
		} else if Data < currentLeaf.Data {
			if currentLeaf.Left == nil {
				return false
			}
			currentLeaf = currentLeaf.Left
		} else {
			if currentLeaf.Right == nil {
				return false
			}
			currentLeaf = currentLeaf.Right
		}
	}
	return false
}

func deleteNodeHelper(Root *Leaf, Data int) *Leaf {
	if Root.Data > Data {
		Root.Left = deleteNodeHelper(Root.Left, Data)
	} else if Root.Data < Data {
		Root.Right = deleteNodeHelper(Root.Right, Data)
	} else {
		if Root.Left == nil {
			return Root.Right
		}
		if Root.Right == nil {
			return Root.Left
		}
		RightSmallest := Root.Right
		for RightSmallest.Left != nil {
			RightSmallest = RightSmallest.Left
		}
		RightSmallest.Left = Root.Left
		return Root.Right
	}
	return Root
}

// DeleteNode deletes a node from the BST and rearranges if needed
func (b *BinarySearchTree) DeleteNode(Data int) {
	if b == nil {
		return
	}
	Root := b.Root
	deleteNodeHelper(Root, Data)
}
