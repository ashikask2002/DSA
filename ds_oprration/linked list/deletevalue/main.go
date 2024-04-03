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
	 //list.deleteValue(5)
	list.DeleteHead()
	//list.DeleteEnd()
	list.displayElements()
}
func (l *LinkedList) deleteValue(value int) {
	temp := l.head

	if temp == nil {
		fmt.Println("list is empty ")
		return
	}

	if temp.data == value {
		l.head = temp.next
		return
	}

	for temp.next != nil && temp.next.data != value {
		temp = temp.next
	}

	if temp.next == nil {
		fmt.Println("Value not found in the list")
		return
	}

	temp.next = temp.next.next
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



func(l *LinkedList) DeleteHead(){
	if l.head == nil{
		fmt.Println("no elements")
	}
	
	if l.head !=  nil{
		l.head = l.head.next
	}
}

func(l *LinkedList) DeleteEnd(){
	if l.head == nil {
		fmt.Println("list is empty")
		return
	} 
	 if l.head.next == nil{
		return
	}
	temp := l.head
		for temp.next.next != nil{
			temp = temp.next
		}
		temp.next = nil
	
}