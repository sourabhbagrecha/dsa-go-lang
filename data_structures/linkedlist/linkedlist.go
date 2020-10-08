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

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	if l.length == 0 {
		l.tail = n
	}
	l.length++
}

func (l *linkedList) printList() {
	myListLength := l.length
	toPrint := l.head
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

func (l *linkedList) push(data int) {
	n := node{data: data}
	if l.length == 0 {
		l.head = &n
		l.tail = &n
	} else {
		l.tail.next = &n
		l.tail = &n
	}
	l.length++
}

func (l *linkedList) pop() {
	if l.length == 0 {
		return
	}
	lastSecondNode := l.head
	nodesRemaining := l.length
	for nodesRemaining > 2 {
		fmt.Println(nodesRemaining, lastSecondNode)
		lastSecondNode = lastSecondNode.next
		nodesRemaining--
	}
	lastSecondNode.next = nil
	l.tail = lastSecondNode
	l.length--
}

//CallLinkedList creates one
func CallLinkedList() {
	mylist := linkedList{}
	// mylist.push(493)
	mylist.push(520)
	mylist.push(617)
	node1 := &node{data: 27}
	node2 := &node{data: 90}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.pop()
	mylist.printList()

}
