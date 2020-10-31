package search

import "fmt"

// ValueEqualToIndexValue returns all the values that have the same index value as their own value
func ValueEqualToIndexValue(arr []int) []int {
	results := []int{}
	for pos, val := range arr {
		if val == pos+1 {
			fmt.Println(val, pos+1)
			results = append(results, val)
		}
	}
	return results
}
