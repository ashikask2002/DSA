package main

import "fmt"

func main() {
	arr := []int{2, 1, 2, 3, 5, 3, 4, 5}
	target := 5
	result := LinearSearch(arr, target)
	if len(result) > 0 {
		fmt.Printf("elenent %d found at %d position", target, result)
	} else {
		fmt.Printf("element %d not found", target)
	}
}

func LinearSearch(arr []int, target int) []int {
	var ar []int
	for i := range arr {
		if arr[i] == target {
			ar = append(ar, i)
		}
	}
	return ar
}
