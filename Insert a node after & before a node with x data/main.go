package main

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
	tail *node
}

func (l *linkedlist) addValues(data int) {
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
	l.tail = newNode
}

func (l *linkedlist) InsertBeforeTheValue(data int, position int) {
	newNode := new(node)
	newNode.data = data
	var previous *node
	for current := l.head; current != nil; current = current.next {
		if current.data == position {
			if position == l.head.data {
				newNode.next = l.head
				l.head = newNode
			} else {
				previous.next = newNode
				newNode.next = current
			}
		}
		previous = current
	}
}

func (l *linkedlist) InsertAfterTheValue(data int, position int) {
	newNode := new(node)
	newNode.data = data
	for current := l.head; current != nil; current = current.next {
		if position == current.data {
			if current == l.tail {
				l.tail.next = newNode
				l.tail = newNode
			} else {
				newNode.next = current.next
				current.next = newNode
			}

		}
	}

}

func (l linkedlist) display() {
	for temp := l.head; temp != nil; temp = temp.next {
		print(temp.data, " ")
	}
	println()
}

func main() {
	l := linkedlist{}
	l.addValues(40)
	l.addValues(50)
	l.addValues(60)
	l.addValues(10)
	l.addValues(20)
	l.addValues(30)
	l.display()
	l.InsertBeforeTheValue(600, 40)
	l.display()
	l.InsertAfterTheValue(600, 30)
	l.display()
}
