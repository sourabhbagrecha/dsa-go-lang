package str

// Alternate solution

// IsStringRotationOfAnother checks if c string is a rotation of s string or not
// func IsStringRotationOfAnother(mainString string, subString string) bool {
// 	mainString = mainString + mainString
// 	return strings.Contains(mainString, subString)
// }

// IsStringRotationOfAnother checks if c string is a rotation of s string or not
func IsStringRotationOfAnother(mainString string, subString string) bool {
	mainString = mainString + mainString
	start, current := -1, 0
	for sPos := range mainString {
		if current == len(subString) {
			return true
		}
		if mainString[sPos] == subString[current] {
			if start == -1 {
				start = sPos
				current = 0
			}
			current++
		} else {
			start = -1
			current = 0
		}
	}
	return false
}
