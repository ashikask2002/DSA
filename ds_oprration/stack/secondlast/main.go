package main

import "fmt"

type Stack struct {
	data []int
	size int
}

func main() {
	l := &Stack{}
	l.push(3)
	l.push(1)
	l.push(6)
	l.push(7)
	l.push(9)
	fmt.Println("array before", l.data)
	scnd := l.secondlast()
	fmt.Println("secondlast element ", scnd)
	
}

func (s *Stack) push(a int) {
	s.data = append(s.data, a)
	s.size++
}

func (s *Stack) Pop() int {
	n := len(s.data) - 1
	toremove := s.data[n]
	s.data = s.data[:n]
	return toremove
}

func (s *Stack) secondlast()int{
if len(s.data) <2{
	fmt.Println("not enough values")
}
return s.data[len(s.data)-2]
}



