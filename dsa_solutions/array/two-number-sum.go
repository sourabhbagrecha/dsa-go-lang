package array

import "sort"

type state struct {
	idx   int
	valid bool
}

// TwoNumberSum to find elements which adds up to the target
func TwoNumberSum(arr []int, target int) []int {
	results := []int{}
	traverseState := make(map[int]bool) //required num: bool(available)

	for _, num := range arr {
		if traverseState[num] {
			results = []int{target - num, num}
			return results
		}
		traverseState[target-num] = true
	}
	return results
}

// TwoNumberSumIndexes to find indexes of elements which adds up to the target
func TwoNumberSumIndexes(arr []int, target int) []int {
	results := []int{-1, -1}
	traverseState := make(map[int]state) //required num: index of previous

	for idx, num := range arr {
		if traverseState[num].valid {
			results = []int{traverseState[num].idx, idx}
			return results
		}
		traverseState[target-num] = state{idx: idx, valid: true}
	}
	return results
}

// TwoNumberSumUsingSort to find indexes of elements which adds up to the target by sorting the array
func TwoNumberSumUsingSort(arr []int, target int) []int {
	results := []int{}
	left, right := 0, len(arr)-1
	sort.Ints(arr)
	for left < right {
		leftEle, rightEle := arr[left], arr[right]
		currentSum := leftEle + rightEle
		if currentSum == target {
			return []int{leftEle, rightEle}
		} else if currentSum > target {
			right--
		} else {
			left++
		}
	}
	return results
}
