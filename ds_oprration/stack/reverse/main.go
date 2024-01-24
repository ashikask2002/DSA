package main

import "fmt"

type Stack struct {
	data []int
}

func main() {
  l := &Stack{}
  l.push(4)
  l.push(3)
  l.push(2)
  l.push(1)
  l.display()
  fmt.Println("array before",l.data)
  stack := reverse(l)
  fmt.Println("array after",stack)

}
func (s *Stack) push(a int) {
	s.data = append(s.data, a)
}

func (s *Stack) Pop() int {
	n := len(s.data) - 1
	toremove := s.data[n]
	s.data = s.data[:n]
	return toremove
}

func (s *Stack) display() {
	fmt.Println("items are",s.data)
}

func reverse(s *Stack) Stack{
	revstack := Stack{}

	for len(s.data) >0{
		revstack.push(s.Pop())
	}
	return revstack
}