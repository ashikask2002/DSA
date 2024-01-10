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
	list.append(1)
	list.append(3)
	list.append(2)
	list.append(7)
	list.append(9)
	list.append(2)
	list.RemoveDuplicates()
	list.display()
}
func (l *LinkedList) append(value int) {
	node := &Node{value, nil}

	if l.head == nil {
		l.head = node
		return
	}

	temp := l.head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node

}

func (l *LinkedList) display() {
	temp := l.head
	if temp == nil {
		fmt.Println("list is empty ")
	}
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}

}

func (list *LinkedList) RemoveDuplicates() {
	if list.head == nil {
		return
	}

	temp := list.head
	for temp != nil {
		runner := temp
		for runner.next != nil {
			if runner.next.data == temp.data {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		temp = temp.next
	}
}
