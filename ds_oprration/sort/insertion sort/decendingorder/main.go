package main

import "fmt"

func main() {

	arr := []int{2, 5, 1, 3, 6, 4}
    fmt.Println("array before sorting", arr)
	n := len(arr)
	result := InsertionSort(arr, n)
	
	fmt.Println("array after sorting", result)
}

func InsertionSort(arr []int,n int) []int{
	for i:= 1; i<n ; i++{
		temp := arr[i]
		j := i -1

		for j >= 0 && arr[j] < temp{
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp 
	}
	return arr
}