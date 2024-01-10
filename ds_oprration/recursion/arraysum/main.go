package main

import "fmt"

func main() {
	arr := []int{1, 4, 3, 7, 5, 9}
	index := 0
	result := recursion(arr, index)
	fmt.Printf("result is %d", result)
}

func recursion(arr []int, index int) int {
	if index == len(arr) {
		return 0
	}
	return arr[index] + recursion(arr, index+1)
}
