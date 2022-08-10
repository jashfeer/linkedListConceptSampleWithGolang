package methods

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Linkedlist struct {
	head *Node
	len  int
}

func (l *Linkedlist) Add(value int) {
	newNode := new(Node)
	newNode.value = value
	if l.head == nil {
		l.head = newNode
	} else {
		temp := l.head
		for ; temp.next != nil; temp = temp.next {
		}
		temp.next = newNode
	}

}
func (l *Linkedlist) Remove(value int) {
	var previous *Node
	for current := l.head; current != nil; current = current.next {
		if current.value == value {
			if l.head == current {
				l.head = current.next
				break
			}else{
				previous.next = current.next
				break
			}
		}
		previous = current
	}
}
func (l Linkedlist) Get(value int) {
	for temp := l.head; temp != nil; temp = temp.next {
		if temp.value == value {
			fmt.Println(temp.value)
		}
	}
}
func (l Linkedlist) Display() {

	for temp := l.head; temp != nil; temp = temp.next {
		fmt.Print(temp.value, " ")
	}
	fmt.Println()
}
