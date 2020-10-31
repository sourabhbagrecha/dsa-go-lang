package search

import "fmt"

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
func BinarySearch(sortedArray []int, query int) bool {
	min := 0
	max := len(sortedArray) - 1
	for min <= max {
		mid := (min + max) / 2
		fmt.Println(min, max, mid)
		if sortedArray[mid] == query {
			return true
		} else if sortedArray[mid] < query {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return false
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
