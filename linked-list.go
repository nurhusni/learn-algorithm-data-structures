package main

import "fmt"

// struct is like class
type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

// In Go, methods are defined outside the struct
// (l *linkedList) is a method receiver
// This indicates that the method is for linked list

// Add a node to the first index
func (l *linkedList) prepend (n *node) {
	// Create a temp variable for the head node
	second := l.head

	// Assign a new node to the head
	l.head = n

	// Link the nodes
	l.head.next = second

	// Increase the length
	l.length++
}

func main() {
	myList := linkedList{}

	node1 := &node{data:12}
	myList.prepend(node1)

	node2 := &node{data:65}
	myList.prepend(node2)

	node3 := &node{data:32}
	myList.prepend(node3)

	fmt.Println(myList)
}