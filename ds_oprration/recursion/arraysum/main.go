package main

import "fmt"

func main() {
	arr := []int{2, 3, 5, 4, 1}
	index := 0
	result := Recursion(arr, index)
	fmt.Println("sum of array is", result)

}

func Recursion(arr []int, index int) int {
	if index == len(arr) {
		return 0
	}

	return arr[index] + Recursion(arr, index+1)
}
