package main

import "fmt"

type Stack struct {
	data []int
}

func main() {
    Mystack := Stack{}

	Mystack.Push(10)
	Mystack.Push(20)
	Mystack.Push(45)
	Mystack.Push(51)
	Mystack.Pop()
    Mystack.display()

}

func (s *Stack) Push(a int) {
	s.data = append(s.data, a)
}

func (s *Stack) Pop()int {
	n := len(s.data) - 1
	toremove := s.data[n]
	s.data = s.data[:n]
	return toremove
}

func (s *Stack) display() {
	fmt.Println("items are ", s.data)
}

func (s *Stack) Topelement() {
	Topelement := s.data[len(s.data)-1]
	fmt.Println("top element is ",Topelement)
}

