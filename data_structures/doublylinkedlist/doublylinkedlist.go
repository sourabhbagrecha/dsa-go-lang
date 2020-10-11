package doublylinkedlist

import (
	"fmt"
	"math"
)

type node struct {
	data int
	prev *node
	next *node
}

//DoublyLinkedList Struct
type DoublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

func (dl *DoublyLinkedList) push(data int) *node {
	newNode := &node{data: data}
	if dl.length == 0 {
		dl.head = newNode
		dl.tail = newNode
	} else {
		dl.tail.next = newNode
		newNode.prev = dl.tail
		dl.tail = newNode
	}
	dl.length++
	return newNode
}

func (dl *DoublyLinkedList) print() {
	listLength := dl.length
	currentNode := dl.head
	for listLength > 0 {
		fmt.Println(currentNode)
		currentNode = currentNode.next
		listLength--
	}
	return
}

func (dl *DoublyLinkedList) pop() {
	if dl.length == 0 {
		return
	}
	if dl.length == 1 {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.tail.prev.next = nil
		dl.tail = dl.tail.prev
	}
	dl.length--
	return
}

func (dl *DoublyLinkedList) shift() {
	if dl.length == 0 {
		return
	}
	if dl.length == 1 {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.head.next.prev = nil
		dl.head = dl.head.next
	}
	dl.length--
	return
}

func (dl *DoublyLinkedList) unshift(data int) {
	newNode := &node{data: data}
	if dl.length == 0 {
		dl.head = newNode
		dl.tail = newNode
	} else {
		dl.head.prev = newNode
		newNode.next = dl.head
		dl.head = newNode
	}
	dl.length++
	return
}

func (dl *DoublyLinkedList) get(position int) *node {
	if position >= dl.length || dl.length == 0 || position < 0 {
		return nil
	}
	middle := int(math.Floor(float64(dl.length) / 2))
	if position <= middle {
		currentNode := dl.head
		for i := 0; i <= middle; i++ {
			if position == i {
				return currentNode
			}
			currentNode = currentNode.next
		}
	} else {
		currentNode := dl.tail
		for i := dl.length - 1; i > middle; i-- {
			if position == i {
				return currentNode
			}
			currentNode = currentNode.prev
		}
	}
	return nil
}

func (dl *DoublyLinkedList) set(position int, data int) *node {
	if position >= dl.length || dl.length == 0 || position < 0 {
		return nil
	}
	nodeToBeUpdated := dl.get(position)
	nodeToBeUpdated.data = data
	return nodeToBeUpdated
}

func (dl *DoublyLinkedList) insert(position int, data int) *node {
	newNode := &node{data: data}
	if position > dl.length || position < 0 {
		return nil
	}
	if position == dl.length {
		return dl.push(data)
	}
	if dl.length == 0 {
		dl.head = newNode
		dl.tail = newNode
	} else {
		originalNode := dl.get(position)
		prevNode := originalNode.prev
		prevNode.next = newNode
		newNode.next = originalNode
		newNode.prev = prevNode
		originalNode.prev = newNode
	}
	dl.length++
	return newNode
}

func (dl *DoublyLinkedList) remove(position int) *node {
	if position > dl.length || dl.length == 0 {
		return nil
	}
	nodeToBeDeleted := dl.get(position)
	if nodeToBeDeleted.next != nil {
		nodeToBeDeleted.next.prev = nodeToBeDeleted.prev
	}
	if nodeToBeDeleted.prev != nil {
		nodeToBeDeleted.prev.next = nodeToBeDeleted.next
	}
	dl.length--
	return nodeToBeDeleted
}

func (dl *DoublyLinkedList) reverse() {
	if dl.length <= 1 {
		return
	}
	listLength := dl.length
	currentNode := dl.head
	for listLength > 0 {
		temp := currentNode.next
		currentNode.next = currentNode.prev
		currentNode.prev = temp
		currentNode = temp
		listLength--
	}
	tempHead := dl.head
	dl.head = dl.tail
	dl.tail = tempHead
	return
}

//CallDoublyLinkedList triggers
func CallDoublyLinkedList() {
	var dl DoublyLinkedList

	dl.pop()
	dl.shift()
	dl.unshift(4)
	dl.unshift(3)
	dl.unshift(2)
	dl.unshift(1)
	dl.unshift(0)
	// dl.set(0, 44444)
	dl.push(77)
	dl.insert(5, 65)
	dl.remove(3)
	dl.print()
	fmt.Println("==================")
	dl.reverse()
	dl.print()
}
