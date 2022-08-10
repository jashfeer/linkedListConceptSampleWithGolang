package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head *node
	tail *node
	len int
}

func (l *linkedList) addLinkedlist(data int) {
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
	l.tail=newNode
}


func (l *linkedList)addFirst(data int){
	newNode:=new(node)
	newNode.data=data
	newNode.next=l.head
	l.head=newNode
}

func(l *linkedList)addLast(data int){
	newNode:=new(node)
	newNode.data=data
	l.tail.next=newNode
	l.tail=newNode
}




func (l linkedList) display() {
	for temp := l.head; temp != nil; temp = temp.next {
		fmt.Print(temp.data," ")
	}
	fmt.Println()
}

func main() {
	l := linkedList{}
	l.addLinkedlist(30)
	l.addLinkedlist(25)
	l.addLinkedlist(50)
	l.addLinkedlist(40)
	l.display()
	fmt.Println(l.head.data)
	fmt.Println(l.tail.data)
	l.addFirst(4)
	l.display()
	fmt.Println(l.head.data)
	fmt.Println(l.tail.data)
	l.addLast(49)
	l.display()
	fmt.Println(l.head.data)
	fmt.Println(l.tail.data)
}
