package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 4, 3, 6}
	fmt.Println("array before sorting", arr)
	n := len(arr)

	result := Selectionsort(arr, n)
	fmt.Println("array after sorting ", result)
}

func Selectionsort(arr []int,n int)[]int{
	for i := 0; i< n-1;i++{
		min := i
		for j := i+1;j<n-1;j++{
			if arr[j]<arr[min]{
				min = j
			}
			arr[i],arr[min] = arr[min],arr[i]
		}
		
	}
	return arr
}