package matrix

// SearchElement checks if the targetis present in the given matrix or not
func SearchElement(mat [][]int, target int) bool {
	m := len(mat)
	n := len(mat[0])
	l, r := 0, m*n-1
	for l < r {
		mid := (l + r - 1) >> 1
		currentElement := mat[mid/m][mid%n]
		if currentElement == target {
			return true
		} else {
			if currentElement < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return false
}
