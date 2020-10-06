package sort

import (
	"fmt"
	"math"
)

func swap(nums []int, firstPosition int, secondPosition int) {
	temp := nums[firstPosition]
	nums[firstPosition] = nums[secondPosition]
	nums[secondPosition] = temp
}

// Bubble sort
func Bubble(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		noSwaps := true
		for j := 0; j < len(nums)-1; j++ {
			fmt.Println(nums[j], nums[j+1])
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
				noSwaps = false
			}
		}
		if noSwaps {
			break
		}
	}
	return nums
}

//Selection sort
func Selection(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minPosition := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minPosition] {
				minPosition = j
			}
		}
		if i != minPosition {
			swap(nums, minPosition, i)
		}
	}
	return nums
}

// Insertion sort
func Insertion(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		currentVal := nums[i]
		var j int
		for j = i - 1; j >= 0 && nums[j] > currentVal; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = currentVal
	}
	return nums
}

func mergeSortedArrays(first []int, second []int) []int {
	nums := []int{}
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			nums = append(nums, first[i])
			i++
		} else {
			nums = append(nums, second[j])
			j++
		}
	}
	for i < len(first) {
		nums = append(nums, first[i])
		i++
	}
	for j < len(second) {
		nums = append(nums, second[j])
		j++
	}
	return nums
}

// Merge sort
func Merge(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	midCalc := float64(len(nums)) / 2.0
	mid := int(math.Floor(midCalc))
	left := Merge(nums[:mid])
	right := Merge(nums[mid:])
	return mergeSortedArrays(left, right)
}
