package array

import "fmt"

// UnionAndIntersectionSorted finds the union and intersection between two arrays
func UnionAndIntersectionSorted(arr, brr []int) (union, intersection []int) {
	union = []int{}
	intersection = []int{}
	var a, b int = 0, 0
	for a < len(arr) && b < len(brr) {
		if arr[a] == brr[b] {
			intersection = append(intersection, arr[a])
			if len(union) == 0 {
				union = append(union, arr[a])
			} else if union[len(union)-1] != arr[a] {
				union = append(union, arr[a])
			}
			a++
			b++
		} else {
			if arr[a] < brr[b] {
				if len(union) == 0 {
					union = append(union, arr[a])
				} else if union[len(union)-1] != arr[a] {
					union = append(union, arr[a])
				}
				a++
			} else {
				if len(union) == 0 {
					union = append(union, arr[a])
				} else if union[len(union)-1] != brr[b] {
					union = append(union, brr[b])
				}
				b++
			}
		}
		fmt.Println(intersection)
	}
	return union, intersection
}
