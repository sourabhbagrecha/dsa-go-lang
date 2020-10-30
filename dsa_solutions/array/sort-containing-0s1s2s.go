package array

// SortContaining0s1s2s Sorts an array that only contains 0s, 1s and 2s
func SortContaining0s1s2s(arr []int) []int {
	lo, mid, hi := 0, 0, len(arr)-1
	for mid <= hi {
		switch arr[mid] {
		case 0:
			swap(arr, lo, mid)
			lo++
			break
		case 1:
			break
		case 2:
			swap(arr, mid, hi)
			hi--
			break
		default:
			break
		}
		mid++
	}
	return arr
}
