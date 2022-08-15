/*
👋 Welcome Gophers! In this challenge, you will be implementing some of the basic functionality of a singly linked-list in Go! 💪

In this challenge, we will be attempted to implement the Add function of a singly linked list.

Challenge Requirements
In this challenge, you will be implementing a very simple, singly linked list that allows you to append new string values to the end of our linked list and ensure that you update the Tail and Head values appropriately.
*/

package main

import "fmt"

type Element struct {
	Value string
	Next  *Element
}

type SinglyLinkedList struct {
	Count int
	Head  *Element
	Tail  *Element
}

func (l *SinglyLinkedList) Append(value string) {
	node := &Element{Value: value}
	// No nodes in the list
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		l.Count++
		// 1 or more nodes in the list
	} else {
		l.Tail.Next = node
		l.Tail = node
		l.Count++
	}
}

// You will have to ensure when you add new elements
// that this method still returns the correct value
func (l *SinglyLinkedList) Size() int {
	return l.Count
}

func (l *SinglyLinkedList) Print() {
	current := l.Head
	fmt.Printf("%+v\n", current.Value)
	for current.Next != nil {
		fmt.Printf("%+v\n", current.Value)
		current = current.Next
	}
}

func main() {
	fmt.Println("Singly Linked List Challenge")

	var llist SinglyLinkedList

	values := []string{"First", "Second", "Third"}
	for _, value := range values {
		llist.Append(value)
	}
	llist.Print()
}
