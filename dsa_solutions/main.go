package main

import (
	"github.com/sourabhbagrecha/first_project/dsa_solutions/matrix"
)

func main() {
	// arr := []int{1, 2, 2, 0, 0, 1, 2}
	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	// fmt.Println(array.Reverse(arr))
	// fmt.Println(array.FindMax(arr))
	// fmt.Println(array.FindMinMax(arr))
	// array.KthMax(arr)
	// fmt.Println(array.RemoveDuplicatesSorted(arr))
	// fmt.Println(array.SortContaining0s1s2s(arr))
	matrix.SpiralTraversal(mat)
}
