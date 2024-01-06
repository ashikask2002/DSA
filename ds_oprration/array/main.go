// add the elements to array
package main

import "fmt"

func main() {
	arr := make([]int, 20)
	fmt.Println("enter the values to array")
	size := 5
	addarray(arr, size)
	displayarray(arr, size)
}
func addarray(arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
}

func displayarray(arr []int, size int) {
	fmt.Println("the array is ")
	for i := 0; i < size; i++ {
		fmt.Printf("%d", arr[i])
	}
}
