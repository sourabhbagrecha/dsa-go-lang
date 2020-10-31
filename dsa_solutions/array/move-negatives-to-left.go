package array

// MoveNegativesToLeft moves all the negative values to the left side of the array
func MoveNegativesToLeft(arr []int) []int {
	origin := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			swap(arr, i, origin)
			origin++
		}
	}
	return arr
}
