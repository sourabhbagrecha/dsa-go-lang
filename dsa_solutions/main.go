package main

import (
	"fmt"

	"github.com/sourabhbagrecha/first_project/dsa_solutions/search"
)

func main() {
	arr := []int{1, 3, 5, 5, 5, 5, 7, 123, 125}
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
	fmt.Println(search.FindFirstAndLastOccurrence(arr, 11))
}
