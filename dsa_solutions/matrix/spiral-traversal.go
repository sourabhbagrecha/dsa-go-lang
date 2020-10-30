package matrix

import (
	"fmt"
)

// SpiralTraversal of the given matrix
func SpiralTraversal(arr [][]int, params ...int) {
	var i, j, m, n int
	if len(params) > 0 {
		i, j, m, n = params[0], params[1], params[2], params[3]
	} else {
		i, j, m, n = 0, 0, len(arr), len(arr[0])
	}
	// If i or j lies outside the matrix
	if i >= m || j >= n {
		return
	}

	// Print First Row
	for p := i; p < n; p++ {
		fmt.Printf("%d ", arr[i][p])
	}

	// Print Last Column
	for p := i + 1; p < m; p++ {
		fmt.Printf("%d ", arr[p][n-1])
	}

	// Print Last Row, if Last and
	// First Row are not same
	if (m - 1) != i {
		for p := n - 2; p >= j; p-- {
			fmt.Printf("%d ", arr[m-1][p])
		}
	}

	// Print First Column,  if Last and
	// First Column are not same
	if (n - 1) != j {
		for p := m - 2; p > i; p-- {
			fmt.Printf("%d ", arr[p][j])
		}
	}

	SpiralTraversal(arr, i+1, j+1, m-1, n-1)
}
