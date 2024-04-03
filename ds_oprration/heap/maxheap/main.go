package main

import "fmt"

type maxheap struct {
	arr []int
}

func (h *maxheap) insert(data int) {
	h.arr = append(h.arr, data)
	h.Buildheap()
}

func (h *maxheap) Buildheap() {
	n := len(h.arr)

	for i := n/2 - 1; i >= 0; i-- {
		h.maxheapify(i)
	}
}

func (h *maxheap) maxheapify(Index int) {
	n := len(h.arr) - 1
	var k int
	for {
		l := leftchild(Index)
		r := rightchild(Index)
		if l > n {
			break
		}
		if r <= n && h.arr[r] > h.arr[l] {
			k = r
		} else {
			k = l
		}
		if h.arr[k] > h.arr[Index] {
			swap(h.arr, k, Index)
			Index = k
		} else {
			break
		}
	}

}

func (h *maxheap) Delete() int {
	if len(h.arr) == 0 {
		return -1
	}
	maxvalue := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.maxheapify(0)
	return maxvalue
}

func leftchild(i int) int {
	return 2*i + 1
}

func rightchild(i int) int {
	return 2*i + 2
}
func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func display(arr []int)	 {
	for _, v := range arr {

		fmt.Print(v, " ")
	}
	fmt.Println(" ")
}

func main() {
	var arr = []int{11, 2, 7, 4, 12, 5, 9, 10}
	h := &maxheap{}
	for _, v := range arr {
		h.insert(v)
	}
	fmt.Println("array is ", arr)
	h.Buildheap()
	fmt.Print("heap")
	display(h.arr)
	h.Delete()
	fmt.Println("after delete")
	display(h.arr)
}
