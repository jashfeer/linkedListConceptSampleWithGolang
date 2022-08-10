package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
}

func (l *linkedlist) addValues(data int) {
	newNode := new(node)
	newNode.data = data
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for ; current.next != nil; current = current.next {
		}
		current.next = newNode
	}

}

func (l linkedlist) ReverseLinkedList(current *node) {
	if current == nil {
		return
	}
	l.ReverseLinkedList(current.next)
	fmt.Printf("%d ", current.data)
}



func (l linkedlist) display() {
	for current := l.head; current != nil; current = current.next {
		print(current.data, " ")
	}
	fmt.Println()
	
}

func main() {
	l := linkedlist{}
	l.addValues(20)
	l.addValues(27)
	l.addValues(60)
	l.addValues(50)
	l.addValues(37)
	l.addValues(10)
	l.display()
	l.ReverseLinkedList(l.head)

}
