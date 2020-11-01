package array

// MaxContiguousSum finds the maximum contigous sum in an array
func MaxContiguousSum(arr []int) int {
	maxSoFar := 0
	maxEndingCurrently := 0
	for i := 0; i < len(arr); i++ {
		maxEndingCurrently += arr[i]
		if maxEndingCurrently > maxSoFar {
			maxSoFar = maxEndingCurrently
		}
		if maxEndingCurrently < 0 {
			maxEndingCurrently = 0
		}
	}
	return maxSoFar
}
