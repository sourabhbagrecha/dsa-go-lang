package main

import (
	// "github.com/sourabhbagrecha/first_project/data_structures/doublylinkedlist"
	// "github.com/sourabhbagrecha/first_project/data_structures/linkedlist"
	"fmt"

	"github.com/sourabhbagrecha/first_project/data_structures/stack"
)

func main() {
	// linkedlist.CallLinkedList()
	// fmt.Println(linkedlist.LinkedList{})
	// doublylinkedlist.CallDoublyLinkedList()
	newStack := stack.Stack{}
	fmt.Println(newStack.Push(40))
	fmt.Println(newStack.Push(140))
	fmt.Println(newStack.Push(220))
	fmt.Println(newStack.Push(642))
	fmt.Println(newStack.Push(790))
	newStack.Pop()
	newStack.Pop()
	newStack.Print()
}
