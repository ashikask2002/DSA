package main

import "fmt"

func main() {
	arr := []int{1, 2, 6, 3, 9, 4}
	target := 6
	result := LinearSearch(arr, target)

	if result != -1 {
		fmt.Printf("element %d found at %d  index", target, result)
	} else {
		fmt.Println("element %d not found",target)
	}
}

func LinearSearch(arr []int, target int) int {
	for i := range arr {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
