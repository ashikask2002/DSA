package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 9, 21, 32, 35}
	target := 4

	result := Binarysearch(arr, target)
	if result != -1 {
		fmt.Printf("element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("element not found\n")
	}

}

func Binarysearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}
