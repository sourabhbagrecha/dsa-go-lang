package search

// FindFirstAndLastOccurrence finds First and last occurrences of the given num
func FindFirstAndLastOccurrence(arr []int, num int) (int, int) {
	i := 0
	start, end := -1, -1
	for i < len(arr) {
		if arr[i] == num {
			end = i
			if start < 0 {
				start = i
			}
		}
		i++
	}
	return start, end
}
