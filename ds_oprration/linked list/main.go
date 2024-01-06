package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	list := LinkedList{}
	list.addHead(1)
	list.addHead(5)
	list.addHead(3)
	list.addHead(9)
	list.displayElements()

}

func (l *LinkedList) addHead(value int) {
	node := &Node{value, nil}

	if l.head == nil {
		l.head = node

	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *LinkedList) displayElements() {
	temp := l.head

	if temp == nil {
		fmt.Println("list is empty")
	} else {
		for temp != nil {
			fmt.Println(temp.data)
			temp = temp.next
		}
	}
}
