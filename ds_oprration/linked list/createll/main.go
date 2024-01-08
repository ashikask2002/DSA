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
	list:= LinkedList{}
	list.append(2)
	list.append(1)
	list.append(8)
	list.append(6)
	list.append(4)
	list.DisplayList()

}

func (link *LinkedList) append(data int) {
	node := &Node{data, nil}

	if link.head == nil {
		link.head = node
		return
	}

	temp := link.head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node

}
func(l *LinkedList) DisplayList(){
	temp :=  l.head

	for temp != nil{
		fmt.Println("%d",temp.data)
		temp = temp.next
	}
}
