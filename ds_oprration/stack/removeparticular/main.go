 package main

import "fmt"

type Stack struct {
	data []int
	size int
}

func main() {
	l := &Stack{}
	l.Push(8)
	l.Push(2)
	l.Push(6)
	l.Push(2)
	l.Push(5)
	fmt.Println("array before ", l.data)
	value := 8
	l.removeElement(value)
	fmt.Println("array after", l.data)
}

func (s *Stack) Push(a int) {
	s.data = append(s.data, a)
	s.size++
}

func (s *Stack) pop() int {
	if s.size == 0 {
		return -1 
	}
	n := len(s.data) - 1
	toremove := s.data[n]
	s.data = s.data[:n]
	s.size--
	return toremove
}

func (s *Stack)removeElement(a int){
	var nestack Stack

	for s.size > 0{
       temp := s.pop()
	   if temp != a{
		nestack.Push(temp)
	   }
	}
	for nestack.size > 0{
		s.Push(nestack.pop())
	}
}