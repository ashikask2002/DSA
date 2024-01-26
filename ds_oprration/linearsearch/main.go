package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 4, 7, 3}
	target := 4
	result := LinearSearch(arr, target)

	if result != -1 {
		fmt.Printf("element %d found at target %d\n", target, result)
	} else {
		fmt.Println("element %d nt found", target)
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
