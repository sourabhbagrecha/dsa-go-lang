package search

// FindInSortedAndRotatedArray checks if the target value is available in the sorted and rotated array or not
func FindInSortedAndRotatedArray(arr []int, target int) bool {
	n := len(arr)
	lo, hi := 0, n-1

	// To find the index with lowest value
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid] > arr[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	pivot := lo
	lo, hi = 0, n-1

	// perform the usual binary search, but with a twist in the real middle element
	for lo <= hi {
		mid := (lo + hi) >> 1
		realMid := (mid + pivot) % n
		if arr[realMid] == target {
			return true
		} else if arr[realMid] > arr[lo] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}
