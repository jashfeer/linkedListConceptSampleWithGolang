package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
	tail *node
	len  int
}

func (l *linkedlist) addLinkedlist(data int) {
	newNode := new(node)
	newNode.data = data
	if l.head == nil {
		l.head = newNode
	} else {
		temp := l.head
		for ; temp.next != nil; temp = temp.next {
		}
		temp.next = newNode
	}

}
func (l *linkedlist) delete(data int) {
	var previous *node
	ok := false
	for temp := l.head; temp != nil; temp = temp.next {
		if temp.data == data {
			if temp == l.head {
				l.head = l.head.next
			} else {
				previous.next = temp.next
			}
			temp=l.head
			ok = true
		}
		previous = temp
	}
	if ok == false {
		fmt.Println("no data")
	}
}

func (l linkedlist) display() {
	for temp := l.head; temp != nil; temp = temp.next {
		fmt.Print(temp.data, " ")
	}
	fmt.Println()
}

func main() {
	l := linkedlist{}
	l.addLinkedlist(2)
	l.addLinkedlist(30)
	l.addLinkedlist(30)
	l.addLinkedlist(2)
	l.addLinkedlist(20)
	l.addLinkedlist(36)
	l.display()
	l.delete(30)
	l.display()
}
