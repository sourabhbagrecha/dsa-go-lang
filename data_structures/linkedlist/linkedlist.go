package linkedlist

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func (l *linkedList) unshift(data int) *node {
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.length++
	return newNode
}

func (l *linkedList) printList() {
	myListLength := l.length
	toPrint := l.head
	fmt.Println("LinkedList:>")
	for myListLength > 0 {
		fmt.Println(toPrint)
		toPrint = toPrint.next
		myListLength--
	}
}

func (l *linkedList) deleteNode(data int) {
	myListLength := l.length
	nodeSelected := l.head
	if nodeSelected.data == data {
		l.head = nodeSelected.next
		l.length--
		return
	}
	for myListLength > 0 {
		if nodeSelected.next.data == data {
			nodeSelected.next = nodeSelected.next.next
			l.length--
			return
		}
		myListLength--
	}
	return
}

func (l *linkedList) push(data int) *node {
	newNode := &node{data: data}
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.length++
	return newNode
}

func (l *linkedList) pop() {
	if l.head == nil {
		return
	}
	current := l.head
	var newTail *node
	for current.next != nil {
		newTail = current
		current = current.next
	}
	current.next = nil
	l.tail = newTail
	l.tail.next = nil
	l.length--
	if l.length == 0 {
		l.head = nil
		l.tail = nil
	}
	return
}

func (l *linkedList) shift() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	l.length--
	if l.length == 0 {
		l.tail = nil
	}
	return
}

func (l *linkedList) get(position int) *node {
	myListLength := l.length
	if position > myListLength || myListLength == 0 {
		return nil
	}
	currentNode := l.head
	for myListLength > 0 {
		if (l.length - myListLength) == position {
			return currentNode
		}
		currentNode = currentNode.next
		myListLength--
	}
	return currentNode
}

func (l *linkedList) set(position int, data int) int {
	myListLength := l.length
	if position > myListLength || myListLength == 0 {
		return -1
	}
	currentNode := l.head
	for myListLength > 0 {
		if (l.length - myListLength) == position {
			currentNode.data = data
			return currentNode.data
		}
		currentNode = currentNode.next
		myListLength--
	}
	return -1
}

func (l *linkedList) insert(position int, data int) *node {
	myListLength := l.length
	newNode := &node{data: data}
	if position > myListLength || myListLength == 0 {
		return nil
	} else if position == myListLength {
		return l.push(data)
	} else if position == 0 {
		return l.unshift(data)
	} else {
		prevNode := l.get(position - 1)
		newNode.next = prevNode.next
		prevNode.next = newNode
		l.length++
		return newNode
	}
}

func (l *linkedList) remove(position int) bool {
	myListLength := l.length
	if position > myListLength || myListLength == 0 {
		return false
	} else if position == 0 {
		l.head = l.head.next
		l.length--
		return true
	} else if position == l.length-1 {
		l.pop()
		return true
	} else {
		prevNode := l.get(position - 1)
		prevNode.next = prevNode.next.next
		l.length--
		return true
	}
}

func (l *linkedList) reverse() {
	if l.length <= 1 {
		return
	}
	myListLength := l.length
	currentNode := l.head
	l.head = l.tail
	l.tail = currentNode
	var next *node
	var prev *node
	for myListLength > 0 {
		next = currentNode.next
		currentNode.next = prev
		prev = currentNode
		currentNode = next
		myListLength--
	}
}

//CallLinkedList creates one
func CallLinkedList() {
	mylist := linkedList{}
	// mylist.push(493)
	mylist.push(520)
	mylist.push(617)
	mylist.unshift(27)
	mylist.unshift(99)
	// mylist.pop()
	mylist.shift()
	mylist.unshift(45)
	mylist.set(2, 333)
	mylist.insert(0, 9211)
	mylist.printList()
	mylist.remove(1)
	mylist.printList()
	mylist.reverse()
	mylist.printList()
}
