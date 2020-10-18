package binaryheap

import (
	"fmt"
	"math"
)

//MaxBinaryHeap struct
type MaxBinaryHeap struct {
	values []int
}

func swap(nums []int, firstPosition int, secondPosition int) {
	temp := nums[firstPosition]
	nums[firstPosition] = nums[secondPosition]
	nums[secondPosition] = temp
}

func (maxHeap *MaxBinaryHeap) insert(data int) {
	maxHeap.values = append(maxHeap.values, data)
	maxHeap.bubbleUp()
}

func (maxHeap *MaxBinaryHeap) bubbleUp() {
	idx := len(maxHeap.values) - 1
	parent := int(math.Floor(float64((idx - 1) / 2)))
	for idx > 0 && maxHeap.values[idx] > maxHeap.values[parent] {
		swap(maxHeap.values, idx, parent)
		idx = parent
		parent = int(math.Floor(float64((idx - 1) / 2)))
	}
}

func (maxHeap *MaxBinaryHeap) extractMax() int {
	if len(maxHeap.values) == 0 {
		return 0
	}
	if len(maxHeap.values) == 1 {
		element := maxHeap.values[0]
		return element
	}
	max := maxHeap.values[0]
	lastElementIdx := len(maxHeap.values) - 1
	swap(maxHeap.values, 0, lastElementIdx)
	maxHeap.values = maxHeap.values[:lastElementIdx]
	maxHeap.sinkDown()
	return max
}

func (maxHeap *MaxBinaryHeap) sinkDown() {
	lastElementIdx := len(maxHeap.values) - 1
	parentIdx := 0
	for parentIdx <= lastElementIdx {
		leftChildIdx := (parentIdx * 2) + 1
		rightChildIdx := leftChildIdx + 1
		newParentIdx := parentIdx
		maxVal := maxHeap.values[newParentIdx]
		if rightChildIdx <= lastElementIdx {
			if maxHeap.values[rightChildIdx] > maxVal {
				maxVal = maxHeap.values[rightChildIdx]
				newParentIdx = rightChildIdx
			}
		}
		if leftChildIdx <= lastElementIdx {
			if maxHeap.values[leftChildIdx] > maxVal {
				maxVal = maxHeap.values[leftChildIdx]
				newParentIdx = leftChildIdx
			}
		}
		if newParentIdx != parentIdx {
			swap(maxHeap.values, parentIdx, newParentIdx)
			parentIdx = newParentIdx
		} else {
			break
		}
	}
	return
}

//CallBinaryHeap triggers
func CallBinaryHeap() {
	var maxHeap MaxBinaryHeap

	fmt.Println(maxHeap)
	maxHeap.insert(15)
	maxHeap.insert(20)
	maxHeap.insert(7000)
	maxHeap.insert(40)
	maxHeap.insert(50)
	maxHeap.insert(100)
	maxHeap.insert(600)
	maxHeap.insert(700)
	maxHeap.insert(400)
	maxHeap.insert(900)
	fmt.Println(maxHeap)
	fmt.Println(maxHeap.extractMax())
	fmt.Println(maxHeap)
	fmt.Println(maxHeap.extractMax())
	fmt.Println(maxHeap)
	fmt.Println(maxHeap.extractMax())
	fmt.Println(maxHeap)
	fmt.Println(maxHeap.extractMax())
	fmt.Println(maxHeap)
}
