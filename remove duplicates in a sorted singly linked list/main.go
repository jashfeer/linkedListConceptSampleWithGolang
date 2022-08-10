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
	l.len++
	l.tail=newNode
}

func (l linkedlist) display() {
	for current := l.head; current != nil; current = current.next {
		fmt.Print(current.data, " ")
	}
	fmt.Println()
}
func (l *linkedlist) sortLinkedlist() {
	for step := 0; step< l.len; step++ {
		pri :=l.head
		for current :=l.head; current!= nil; current = current.next {
			if pri.data > current.data {
				pri.data, current.data = current.data, pri.data
			}
			pri = current
		}
	}
}

func (l *linkedlist)deletDuplicate(){
	pri :=l.head
	for current :=l.head.next; current!= nil; current = current.next {
		if pri.data==current.data{
			if current==l.tail{
				pri.next=nil
				current=pri
			}else{
				pri.next=current.next
				current=pri
			}
		}
		pri = current
	}
}


func main() {
	l := linkedlist{}
	l.addValues(1)
	l.addValues(4)
	l.addValues(2)
	l.addValues(1)
	l.addValues(3)
	l.addValues(4)
	l.addValues(2)
	l.addValues(3)
	l.display()
	l.sortLinkedlist()
	l.display()
	l.deletDuplicate()
	l.display()
}
