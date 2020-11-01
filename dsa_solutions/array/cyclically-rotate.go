package array

// CyclicallyRotate rotates the array by one position cyclically
func CyclicallyRotate(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	last := arr[len(arr)-1]
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		arr[i] = prev
		prev = temp
	}
	arr[0] = last
	return arr
}
