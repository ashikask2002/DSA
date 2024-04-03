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
	// list.addHead(1)
	// list.addHead(5)
	// list.addHead(3)
	// list.addHead(9)
list.addend(5)
list.addend(3)
list.addend(7)
list.addend(2)

//    list.deleteValue(4)
    
   list.displayElements()

}

func(l *LinkedList) addHead(value int){
	node := &Node{value,nil}

	if l.head == nil{
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

//adde the node to linked list at end


func(l *LinkedList)addend(value int){
	node := &Node{value,nil}
	temp := l.head

	if temp != nil {
      l.head = node
	} else {
		for temp.next != nil{
			temp = temp.next
		}
		temp.next = node
	}
}



func (l *LinkedList) deleteValue(value int){
	temp := l.head
  if temp == nil{
	fmt.Println("list is empty")
	return
  } 

  if temp.data == value {
	l.head = temp.next
	return
  }

  for temp.next != nil {
	if temp.next.data == value{
		temp.next = temp.next.next
		return
	}
	if temp.next != nil{
		temp = temp.next
	}
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


