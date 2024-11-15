package main

import (
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (l *LinkedList[T]) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	initList := LinkedList[int]{}
	initList.Add(1)
	initList.Add(2)
	initList.Add(3)
	initList.Add(4)
	initList.Add(5)
	initList.Display()
}
