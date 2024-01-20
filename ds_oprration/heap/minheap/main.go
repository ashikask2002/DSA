package main

import "fmt"

type minheap struct {
	arr []int
}

func (m *minheap) Insert(value int) {
	m.arr = append(m.arr, value)
	m.buildheap()
}

func (m *minheap) buildheap() {
	n := len(m.arr)
	for i := n/2 - 1; i >= 0; i-- {
		m.minheapify(i)
	}
}

func (m *minheap) minheapify(index int) {
	n := len(m.arr) - 1
	var k int
	for {
		l := leftchild(index)
		r := rightchild(index)
		if l > n {
			break
		}
		if r <= n && m.arr[r] < m.arr[l] {
			k = r
		} else {
			k = l
		}
		if m.arr[k] < m.arr[index] {
			swap(m.arr, k, index)
			index = k
		} else {
			break
		}
	}
}

func (m *minheap) Delete() int {
	if len(m.arr) == 0 {
		return -1
	}

	minvalue := m.arr[0]
	m.arr[0] = m.arr[len(m.arr)-1]
	m.arr = m.arr[:len(m.arr)-1]
    m.minheapify(0)
	return minvalue
}

func leftchild(index int) int {
	return 2*index + 1
}
func rightchild(index int) int {
	return 2*index + 2
}
func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func (m *minheap) display() {
	for _, val := range m.arr {
		fmt.Print(val, " ")
	}
}

func main(){
	arr := []int{4,2,7,5,8,12,10}
	l := &minheap{}
	for _,val := range arr{
		l.Insert(val)
	}
	l.buildheap()
	fmt.Println("heap")
	l.display()

	del := l.Delete()
	fmt.Println("deleted element",del)
	fmt.Println("after delete")
	l.display()
}
