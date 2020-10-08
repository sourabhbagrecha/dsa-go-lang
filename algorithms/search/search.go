package search

import (
	"math"
)

//LinearSearch finds the position of an string in the list
func LinearSearch(names []string, query string) int {
	for i := 0; i < len(names); i++ {
		if query == names[i] {
			return i
		}
	}
	return -1
}

//BinarySearch finds the position of match optimally in the list
func BinarySearch(sortedNames []int, query int) int {
	min := 0
	max := len(sortedNames) - 1
	for min < max {
		cursor := math.Floor((float64(min) + float64(max)) / 2)
		currentVal := sortedNames[int(cursor)]
		if currentVal == query {
			return int(cursor)
		} else if (min + 1) == max {
			min = max
		} else if currentVal > query {
			max = int(cursor)
		} else {
			min = int(cursor)
		}
	}
	return -1
}

// NaiveStringSearch gives the count of matches found of given string in a long string
func NaiveStringSearch(longString string, shortString string) int {
	count := 0
	longRunes := []rune(longString)
	shortRunes := []rune(shortString)

	for i := 0; i < len(longRunes); i++ {
		if longRunes[i] != shortRunes[0] {
			continue
		}
		matchFound := true

		for j := 0; j < len(shortRunes); j++ {
			if i+j > len(longRunes)-1 {
				matchFound = false
				break
			}
			if longRunes[i+j] != shortRunes[j] {
				matchFound = false
				break
			}
		}

		if matchFound {
			count++
		}
	}
	return count
}
