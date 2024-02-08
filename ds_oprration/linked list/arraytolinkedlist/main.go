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
	arr := []int{2, 3, 4, 1, 5}
	fmt.Println(arr)
	list.arraytolinkedList(arr)
	fmt.Println("linked list is")
	list.display()
	
}

func (l *LinkedList) arraytolinkedList(arr []int) {
   for _, i := range arr{
	node := &Node{i,nil}
	if l.head == nil{
		l.head = node
	}else {
	temp := l.head
	for temp.next != nil{
		temp = temp.next
	}
	temp.next = node
   }
}
}

func(l *LinkedList) display(){
	temp := l.head

	if temp == nil{
		fmt.Println("list is empty")
	}else {
		for temp != nil{
			fmt.Printf("%d\n",temp.data)
			temp = temp.next
		}
	}
}

