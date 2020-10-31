package str

// FindDuplicates find the characters that are occuring more than once in a string
func FindDuplicates(s string) map[string]int {
	state := make(map[int32]int)
	results := make(map[string]int)
	for _, c := range s {
		state[c]++
	}
	for char, x := range state {
		if x > 1 {
			results[string(char)] = x
		}
	}
	return results
}
