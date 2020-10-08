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

// pivot helper
func pivot(nums []int, start int, end int) int {
	swapIndex := start
	pivotElement := nums[swapIndex]
	for i := start + 1; i <= end; i++ {
		if nums[i] < pivotElement {
			swapIndex++
			swap(nums, i, swapIndex)
		}
	}
	swap(nums, swapIndex, start)
	return swapIndex
}

// Quick Sort
func Quick(nums []int, options ...int) []int {
	left := 0
	right := len(nums) - 1
	if len(options) > 0 {
		left = options[0]
		right = options[1]
	}
	if left < right {
		fmt.Println(nums, left, right)
		pivotIndex := pivot(nums, left, right)
		fmt.Println(nums, left, right, pivotIndex)
		Quick(nums, left, pivotIndex-1)
		Quick(nums, pivotIndex+1, right)
	}
	return nums
}

func getDigit(num int, positionFromRight int) int {
	if num < 0 {
		num = -1 * num
	}
	divisor := int(math.Pow10(positionFromRight))
	remainder := int(num / divisor)
	digit := remainder % divisor
	return digit
}

func numOfDigits(num int) int {
	if num == 0 {
		return 1
	}
	if num < 0 {
		num = num * -1
	}
	digits := int(math.Floor(math.Log10(float64(num)))) + 1
	return digits
}

func mostDigits(nums []int) int {
	maxDigits := 0
	for i := 0; i < len(nums); i++ {
		maxDigits = int(math.Max(float64(numOfDigits(nums[i])), float64(maxDigits)))
	}
	return maxDigits
}

// Radix sort
func Radix(nums []int) []int {
	maxDigits := mostDigits(nums)
	for k := 0; k <= maxDigits; k++ {
		var digitBuckets [][]int
		for digitBucket := 0; digitBucket < 10; digitBucket++ {
			var emptyBucket []int
			digitBuckets = append(digitBuckets, emptyBucket)
		}
		for i := 0; i < len(nums); i++ {
			digit := getDigit(nums[i], k)
			digitBuckets[digit] = append(digitBuckets[digit], nums[i])
		}
		var newNums []int
		for digitBucket := 0; digitBucket < len(digitBuckets); digitBucket++ {
			newNums = append(newNums, digitBuckets[digitBucket]...)
		}
		nums = newNums
	}
	return nums
}
