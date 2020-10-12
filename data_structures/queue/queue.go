package queue

import "fmt"

type node struct {
	data int
	next *node
}

//Queue structure
type Queue struct {
	first  *node
	last   *node
	length int
}

//Print the complete queue
func (queue *Queue) Print() {
	currentNode := queue.first
	for currentNode != nil {
		fmt.Println(currentNode)
		currentNode = currentNode.next
	}
	return
}

//Enqueue adds a new node at the end of the queue
func (queue *Queue) Enqueue(data int) int {
	newNode := &node{data: data}
	if queue.length == 0 {
		queue.first = newNode
		queue.last = newNode
	} else {
		queue.last.next = newNode
		queue.last = queue.last.next
	}
	queue.length++
	return queue.length
}

//Dequeue adds a new node at the end of the queue
func (queue *Queue) Dequeue() int {
	if queue.length == 0 {
		return 0
	}

	if queue.length == 1 {
		queue.first = nil
		queue.last = nil
	} else {
		queue.first = queue.first.next
	}
	queue.length--
	return queue.length
}
