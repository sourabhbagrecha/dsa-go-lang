package main

import (
	"fmt"

	"github.com/sourabhbagrecha/first_project/data_structures/graph"
)

// "github.com/sourabhbagrecha/first_project/data_structures/binarysearchtree"

func main() {
	// linkedlist.CallLinkedList()
	// fmt.Println(linkedlist.LinkedList{})
	// doublylinkedlist.CallDoublyLinkedList()

	//Stack Data Structure

	// newStack := stack.Stack{}
	// fmt.Println(newStack.Push(40))
	// fmt.Println(newStack.Push(140))
	// fmt.Println(newStack.Push(220))
	// fmt.Println(newStack.Push(642))
	// fmt.Println(newStack.Push(790))
	// newStack.Pop()
	// newStack.Pop()
	// newStack.Print()

	//Queue Data Structure

	// newQueue := queue.Queue{}
	// newQueue.Enqueue(70)
	// newQueue.Enqueue(700)
	// newQueue.Enqueue(800)
	// newQueue.Enqueue(9900)
	// newQueue.Enqueue(11100)
	// newQueue.Dequeue()
	// newQueue.Dequeue()
	// newQueue.Print()

	//Binary Search Tree Data Structure

	// newTree := binarysearchtree.BinarySearchTree{}
	// newTree.Insert(40)
	// newTree.Insert(400)
	// newTree.Insert(49)
	// newTree.Insert(11)
	// newTree.Insert(10)
	// newTree.Insert(13)
	// newTree.Insert(19)
	// newTree.Insert(428)
	// newTree.Print()
	// dataToFind := 4000
	// foundString := "NOT"
	// if newTree.Find(dataToFind) {
	// 	foundString = ""
	// }
	// fmt.Println(dataToFind, "is", foundString, "available")
	// fmt.Println(newTree.BFSTraversal())
	// fmt.Println("Pre Order Traversal", newTree.DFSTraversal("pre-order"))
	// fmt.Println("In Order Traversal", newTree.DFSTraversal("in-order"))
	// fmt.Println("Post Order Traversal", newTree.DFSTraversal("post-order"))

	// binaryheap.CallBinaryHeap()

	// hashTable := hashtable.HashTable{}
	// hashTable.Set("Sourabh", "Bagrecha")
	// hashTable.Set("Jeff", "Bezos")
	// hashTable.Set("John", "Parker")
	// fmt.Println(hashTable.Get("Akansha"))
	// fmt.Println("Results: ", hashTable.GetAllKeys())
	// fmt.Println("Results: ", hashTable.GetAllValues())

	graph := graph.Graph{}
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddVertexWithEdges("F", []string{"E", "A"})
	graph.AddEdge("A", "C")
	fmt.Println(graph)
	graph.RemoveVertex("F")
	// graph.RemoveEdge("A", "C")
	fmt.Println(graph)
}