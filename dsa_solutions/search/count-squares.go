package search

import (
	"math"
)

// CountSquares returns the possible perfect square less than the given num
func CountSquares(num int) int {
	if num < 1 {
		return 0
	}
	// count := 0
	// for i := 1; int(math.Pow(float64(i), float64(2))) < num; i++ {
	// 	count++
	// }

	return int(math.Sqrt(float64(num - 1)))
}
