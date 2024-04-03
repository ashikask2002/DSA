package main

import "fmt"

func main() {
	arr := []int{5, 1, 8, 3, 9, 2, 5}
	fmt.Println("unsorted array",arr)
	n := len(arr)
	heapsortMax(arr,n)
	fmt.Println("ascending",arr)
	heapsortMin(arr,n)
	fmt.Println("descending",arr)

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

func heapsortMin(arr []int,n int){
	for i := n/2 - 1 ; i>= 0; i--{
		minheapify(arr,n,i)
	}
	for i := n-1 ; i >= 0; i--{
		arr[0],arr[i] = arr[i],arr[0]
		minheapify(arr,i,0)
	}
}

func minheapify(arr[]int,n,i int){
   smallets := i
   l := 2 *i +1
   r := 2 *i +2
   for l <n && arr[l] < arr[smallets]{
	smallets = l
   }
   for r < n && arr[r] < arr[smallets]{
	smallets = r
   } 
   if smallets != i {
	arr[i],arr[smallets] = arr[smallets],arr[i]
	minheapify(arr,n,smallets)
   }
}