package str

import "math"

// IsPalindrome checks if the string is a palindrome(reversed string is equal to the original string) or not
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i := 0; i < int(math.Floor(float64(len(runes)/2))); i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}
