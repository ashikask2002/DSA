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
	list.append(5)
	list.append(9)
	list.deleteValue(5)
	list.displayElements()
}

func (l *LinkedList) append(value int) {
	node := &Node{value, nil}

	if l.head == nil {
		l.head = node
		return
	}

	temp := l.head
	if temp.next != nil {
		temp = temp.next
	}
	temp.next = node

}
func (l *LinkedList) displayElements() {
	temp := l.head
	if temp == nil {
		fmt.Println("list is empty")
	}
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func(l *LinkedList) deleteValue(value int){
	temp := l.head

	if temp == nil {
		fmt.Println("list is empty ")
		return
	}
	if temp.data == value{
		temp = temp.next
		return
	}

	for temp.next.data == value{
		temp.next = temp.next.next
		return
	}
    if temp.next != nil {
		temp = temp.next
	}


}
