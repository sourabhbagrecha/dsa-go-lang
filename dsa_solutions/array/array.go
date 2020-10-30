package array

import (
	"fmt"
	"math"
)

func swap(nums []int, firstPosition int, secondPosition int) {
	temp := nums[firstPosition]
	nums[firstPosition] = nums[secondPosition]
	nums[secondPosition] = temp
}

//Reverse the given array
func Reverse(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		swap(arr, i, len(arr)-i-1)
	}
	return arr
}

//FindMax value in the given array
func FindMax(arr []int) int {
	max := int(math.Inf(-1))
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//FindMin value in the given array
func FindMin(arr []int) int {
	min := 2147483648
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

//FindMinMax returns min and max with least iterations possible
func FindMinMax(arr []int) map[string]int {
	minMax := map[string]int{}
	if len(arr) == 1 {
		minMax["min"] = arr[0]
		minMax["max"] = arr[0]
		return minMax
	}

	if arr[0] > arr[1] {
		minMax["min"] = arr[1]
		minMax["max"] = arr[0]
	} else {
		minMax["min"] = arr[0]
		minMax["max"] = arr[1]
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] > minMax["max"] {
			minMax["max"] = arr[i]
		} else if arr[i] < minMax["min"] {
			minMax["min"] = arr[i]
		}
	}
	return minMax
}

//KthMax finds the kth max element in the given array
func KthMax(arr []int, k int) {
	maxData := map[int]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j <= int(math.Min(float64(k-1), float64(i))); j++ {
			if arr[i] > maxData[j] {

			}
		}
	}
	fmt.Println(maxData)
}
