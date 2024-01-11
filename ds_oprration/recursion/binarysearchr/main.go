package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 8, 9}
	value := 1
	result := Recursion(arr, value, 0, len(arr)-1)
	fmt.Printf("value %d is found at %d", value, result)
}

func Recursion(arr []int, value int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if arr[mid] == value {
		return mid
	} else if arr[mid] < value {
		left = mid + 1
	} else {
		right = mid - 1
	}
	return Recursion(arr, value, left, right)

}
