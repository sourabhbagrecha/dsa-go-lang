package str

import (
	"math"
	"strings"
)

// Reverse the given string
func Reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < int(math.Floor(float64(len(runes)/2))); i++ {
		temp := runes[i]
		runes[i] = runes[len(runes)-i-1]
		runes[len(runes)-i-1] = temp
	}
	reversedStringArr := make([]string, len(runes))
	for i, runeItem := range runes {
		reversedStringArr[i] = string(runeItem)
	}
	return strings.Join(reversedStringArr, "")
}
