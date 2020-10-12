package stack

import "fmt"

type node struct {
	data int
	next *node
}

//Stack data structure
type Stack struct {
	first  *node
	last   *node
	length int
}

//Print the complete stack
func (stack *Stack) Print() {
	currentNode := stack.first
	for currentNode != nil {
		fmt.Println(currentNode)
		currentNode = currentNode.next
	}
	return
}

//Push the data into the stack
func (stack *Stack) Push(data int) int {
	newNode := &node{data: data}
	if stack.length == 0 {
		stack.first = newNode
		stack.last = newNode
	} else {
		temp := stack.first
		stack.first = newNode
		stack.first.next = temp
	}
	stack.length++
	return stack.length
}

//Pop the top node from the stack
func (stack *Stack) Pop() int {
	if stack.length == 0 {
		return 0
	} else if stack.length == 1 {
		stack.first = nil
		stack.last = nil
	} else {
		temp := stack.first
		stack.first = temp.next
	}
	stack.length--
	return stack.length
}
