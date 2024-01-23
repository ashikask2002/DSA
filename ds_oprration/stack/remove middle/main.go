package main

import "fmt"

type Stack struct {
	data []int
	size int
}

func main() {
	l := &Stack{}
	l.Push(5)
	l.Push(3)
	l.Push(7)
	l.Push(2)
	l.Push(4)
	fmt.Println("array before", l.data)
	middleIndex := l.size / 2
	currentIndex := 0
	mid := l.Removemiddle(middleIndex, currentIndex)
	fmt.Println("array after",l.data)
	fmt.Println("middle element is ",mid)
} 

func (s *Stack) Push(a int) {
	s.data = append(s.data, a)
	s.size++

}

func (s *Stack) Pop() int {
	n := len(s.data) - 1
	toremove := s.data[n]
	s.data = s.data[:n]
	return toremove
}

func (s *Stack) Removemiddle(middleIndex int, currentIndex int) int {
	if middleIndex == currentIndex {
		return s.Pop()
	}
	temp := s.Pop()
	mid := s.Removemiddle(middleIndex, currentIndex+1)
	s.Push(temp)
	return mid
}
