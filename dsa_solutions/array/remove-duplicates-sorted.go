package array

// RemoveDuplicatesSorted removes duplicate items from a sorted array
func RemoveDuplicatesSorted(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
	}
	return arr[:i+1]
}
