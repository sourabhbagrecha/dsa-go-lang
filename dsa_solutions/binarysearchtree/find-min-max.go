package binarysearchtree

import (
	bst "github.com/sourabhbagrecha/dsa-go-lang/data_structures/binarysearchtree"
)

// FindMinMax returns the minimum and maximum value in a BST
func FindMinMax(tree *bst.BinarySearchTree) (int, int) {
	minNode, maxNode := tree.Root, tree.Root
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	for maxNode.Right != nil {
		maxNode = maxNode.Right
	}
	return minNode.Data, maxNode.Data
}
