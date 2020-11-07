package main

import (
	"fmt"

	"github.com/sourabhbagrecha/dsa-go-lang/dsa_solutions/array"
)

// bst "github.com/sourabhbagrecha/dsa-go-lang/data_structures/binarysearchtree"

// var newTree bst.BinarySearchTree

// func init() {
// 	// creating demo tree
// 	newTree = bst.BinarySearchTree{}
// 	newTree.Insert(50)
// 	newTree.Insert(40)
// 	newTree.Insert(450)
// 	newTree.Insert(530)
// 	newTree.Insert(150)
// 	newTree.Insert(1500)
// 	newTree.Insert(2500)
// 	fmt.Println("In Order Traversal", newTree.DFSTraversal("in-order"))
// 	fmt.Println(newTree)
// }

func main() {
	// arr := []int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9}
	// brr := []int{-15, 1, 2, 7, 123, 125}
	// mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	// s := "sourabhbaruos"
	// fmt.Println(array.Reverse(arr))
	// fmt.Println(array.FindMax(arr))
	// fmt.Println(array.FindMinMax(arr))
	// array.KthMax(arr)
	// fmt.Println(array.RemoveDuplicatesSorted(arr))
	// fmt.Println(array.SortContaining0s1s2s(arr))
	// matrix.SpiralTraversal(mat)
	// fmt.Println(str.Reverse(s))
	// fmt.Println(str.IsPalindrome(s))
	// fmt.Println(search.FindFirstAndLastOccurrence(arr, 11))
	// fmt.Println(array.MoveNegativesToLeft(arr))
	// fmt.Println(array.UnionAndIntersectionSorted(arr, brr))
	// fmt.Println(matrix.SearchElement(mat, 9))
	// fmt.Println(str.FindDuplicates("sourabhbagrecha"))
	// fmt.Println(search.ValueEqualToIndexValue(arr))
	// fmt.Println(search.FindInSortedAndRotatedArray(arr, 4))
	// fmt.Println(binarysearchtree.FindMinMax(&newTree))
	// fmt.Println(array.CyclicallyRotate(arr))
	// fmt.Println(str.IsStringRotationOfAnother("sourabh", "rabhso"))
	// fmt.Println(search.CountSquares(3))
	// fmt.Println(array.MaxContiguousSum([]int{-2, -3, 4, -1, -2, 1, 5}))
	// fmt.Println(array.MinimizeTheHeights([]int{3, 9, 12, 16, 20}, 3))
	// fmt.Println(array.MinimumJumpsToTheEnd(arr))
	fmt.Println(array.TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 4))
	fmt.Println(array.TwoNumberSumUsingSort([]int{3, 5, -4, 8, 11, 1, -1, 6}, 4))
	fmt.Println(array.TwoNumberSumIndexes([]int{3, 5, -4, 8, 11, 1, -1, 6}, 4))
}
