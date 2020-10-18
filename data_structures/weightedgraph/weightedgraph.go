package weightedgraph

import (
	"fmt"
	"math"
)

var posInfinity int = -1 * (int(math.Inf(1)) + 1)

func swap(nums []distance, firstPosition int, secondPosition int) {
	temp := nums[firstPosition]
	nums[firstPosition] = nums[secondPosition]
	nums[secondPosition] = temp
}

type connection struct {
	node   string
	weight int
}

//WeightedGraph Data Structure
type WeightedGraph struct {
	adjacencyList map[string][]connection
	visited       map[string]bool
}

//AddVertex add a new vertex into the weighted graph
func (w *WeightedGraph) AddVertex(vertex string) WeightedGraph {
	if w.adjacencyList == nil {
		w.adjacencyList = make(map[string][]connection)
	}
	w.adjacencyList[vertex] = []connection{}
	return *w
}

//AddEdge add a new connection between two vertices in the weighted graph
func (w *WeightedGraph) AddEdge(vertex1 string, vertex2 string, weight int) WeightedGraph {
	if w.adjacencyList == nil {
		w.adjacencyList = make(map[string][]connection)
	}
	w.adjacencyList[vertex1] = append(w.adjacencyList[vertex1], connection{vertex2, weight})
	w.adjacencyList[vertex2] = append(w.adjacencyList[vertex2], connection{vertex1, weight})
	return *w
}

type distance struct {
	val      string
	priority int
}

type helperPriorityQueue struct {
	values []distance
}

//Enqueue data into the priority queue
func (h *helperPriorityQueue) enqueue(val string, priority int) helperPriorityQueue {
	newNode := &distance{val, priority}
	h.values = append(h.values, *newNode)
	if len(h.values) > 1 {
		h.bubbleUp()
	}
	return *h
}

func (h *helperPriorityQueue) bubbleUp() {
	idx := len(h.values) - 1
	parentIdx := int((idx - 1) / 2)
	for idx > 0 && h.values[idx].priority < h.values[parentIdx].priority {
		swap(h.values, idx, parentIdx)
		idx = parentIdx
		parentIdx = int((idx - 1) / 2)
	}
}

//Dequeue the element with the lowest value of priority
func (h *helperPriorityQueue) dequeue() distance {
	polledNode := h.values[0]
	swap(h.values, 0, len(h.values)-1)
	h.values = h.values[:len(h.values)-1]
	h.sinkDown()
	return polledNode
}

func (h *helperPriorityQueue) sinkDown() {
	idx := 0
	lastElementIdx := len(h.values) - 1
	for idx <= lastElementIdx {
		newIdx := idx
		leftChildIdx := (idx * 2) + 1
		rightChildIdx := leftChildIdx + 1
		minPriority := h.values[idx].priority
		if leftChildIdx <= lastElementIdx {
			if h.values[leftChildIdx].priority < minPriority {
				newIdx = leftChildIdx
				minPriority = h.values[newIdx].priority
			}
		}
		if rightChildIdx <= lastElementIdx {
			if h.values[rightChildIdx].priority < minPriority {
				newIdx = rightChildIdx
				minPriority = h.values[newIdx].priority
			}
		}
		if idx == newIdx {
			break
		} else {
			swap(h.values, idx, newIdx)
			idx = newIdx
		}
	}
}

//ApplyDijkstras will calculate the shortest distance for every node to every other node
func (w *WeightedGraph) ApplyDijkstras(startVertex string, endVertex string) {
	p := helperPriorityQueue{}
	distances := make(map[string]int)
	previous := make(map[string]string)
	path := []string{}

	for k := range w.adjacencyList {
		var dist int
		if k == startVertex {
			dist = 0
		} else {
			dist = posInfinity //equivalent to positive-infinity
		}
		distances[k] = dist
		previous[k] = ""
		p.enqueue(k, dist)
	}
	for len(p.values) > 0 {
		smallest := p.dequeue().val
		if smallest == endVertex {
			path = append(path, smallest)
			for previous[smallest] != "" {
				newArr := []string{previous[smallest]}
				path = append(newArr, path...)
				smallest = previous[smallest]
			}
			fmt.Println(path)
			break
		}
		if distances[smallest] != posInfinity {
			currentAdj := w.adjacencyList[smallest]
			for i := 0; i < len(currentAdj); i++ {
				nextNode := currentAdj[i]
				// calculate new distance to neighboring node
				candidate := distances[smallest] + nextNode.weight
				if candidate < distances[nextNode.node] {
					distances[nextNode.node] = candidate
					previous[nextNode.node] = smallest
					p.enqueue(nextNode.node, candidate)
				}
			}
		}
	}
}
