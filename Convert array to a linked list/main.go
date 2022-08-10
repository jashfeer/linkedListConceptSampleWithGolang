package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
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


func (l linkedlist)display(){
	for temp:=l.head;temp!=nil;temp=temp.next{
		fmt.Print(" ",temp.data)
	}
	fmt.Println()
}

func main() {
l:=linkedlist{}
arr:=[...] int {2,4,6,2,4,3,5,6,7,8,24,20}
for i:=range arr {
	l.addLinkedlist(arr[i])
}
l.display()

}
