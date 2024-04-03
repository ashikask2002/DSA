package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func main() {
list := DoublyLinkedList{}
list.adddll(1)
list.adddll(7)
list.adddll(3)
list.adddll(8)
list.adddll(2)

// list.deletehead()
// list.deleteend()
list.display()

}

func(l *DoublyLinkedList) adddll(data int){
	node := &Node{data,nil,nil}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
    

}

func (l *DoublyLinkedList) display() {
	temp := l.head

	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func(l *DoublyLinkedList) deletehead(){
    
	if l.head == nil{
		fmt.Println("list is empty")
	} else {
       if l.head.next != nil{
		l.head = l.head.next
		l.head.prev = nil
	   }
	}

}

func(l *DoublyLinkedList) deleteend(){
	if l.head == nil{
		fmt.Println("empty")
		return
	}
	if l.head == l.tail{
		l.head, l.tail = nil,nil
		return 
	}
	l.tail = l.tail.prev
	l.tail.next = nil
}