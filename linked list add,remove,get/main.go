package main

import (
	methods "list/src"
)

func main() {
	l := methods.Linkedlist{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Display()
	l.Remove(10)
	l.Remove(15)
	l.Remove(1)
	l.Display()
	l.Get(15)
}
