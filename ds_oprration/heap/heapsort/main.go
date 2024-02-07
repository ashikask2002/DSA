package main

import "fmt"

func main() {
	arr := []int{5, 1, 8, 3, 0, 2, 5}
	fmt.Println("unsorted array")
	n := len(arr)
	heapsortMax(arr,n)
	fmt.Println("ascending",arr)

}

func heapsortMax(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		maxheapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxheapify(arr, i, 0)
	}
}

func maxheapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	for l < n && arr[l] > arr[largest] {
		largest = l
	}
	for r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxheapify(arr, n, largest)
	}
}
