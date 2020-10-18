package priorityqueue

func swap(nums []node, firstPosition int, secondPosition int) {
	temp := nums[firstPosition]
	nums[firstPosition] = nums[secondPosition]
	nums[secondPosition] = temp
}

type node struct {
	val      string
	priority int
}

//PriorityQueue data structure
type PriorityQueue struct {
	values []node
}

//Enqueue data into the priority queue
func (p *PriorityQueue) Enqueue(val string, priority int) PriorityQueue {
	newNode := &node{val, priority}
	p.values = append(p.values, *newNode)
	if len(p.values) > 1 {
		p.bubbleUp()
	}
	return *p
}

func (p *PriorityQueue) bubbleUp() {
	idx := len(p.values) - 1
	parentIdx := int((idx - 1) / 2)
	for idx > 0 && p.values[idx].priority < p.values[parentIdx].priority {
		swap(p.values, idx, parentIdx)
		idx = parentIdx
		parentIdx = int((idx - 1) / 2)
	}
}

//Dequeue the element with the lowest value of priority
func (p *PriorityQueue) Dequeue() node {
	polledNode := p.values[0]
	swap(p.values, 0, len(p.values)-1)
	p.values = p.values[:len(p.values)-1]
	p.sinkDown()
	return polledNode
}

func (p *PriorityQueue) sinkDown() {
	idx := 0
	lastElementIdx := len(p.values) - 1
	for idx <= lastElementIdx {
		newIdx := idx
		leftChildIdx := (idx * 2) + 1
		rightChildIdx := leftChildIdx + 1
		minPriority := p.values[idx].priority
		if leftChildIdx <= lastElementIdx {
			if p.values[leftChildIdx].priority < minPriority {
				newIdx = leftChildIdx
				minPriority = p.values[newIdx].priority
			}
		}
		if rightChildIdx <= lastElementIdx {
			if p.values[rightChildIdx].priority < minPriority {
				newIdx = rightChildIdx
				minPriority = p.values[newIdx].priority
			}
		}
		if idx == newIdx {
			break
		} else {
			swap(p.values, idx, newIdx)
			idx = newIdx
		}
	}
}
