package main

import "fmt"

func main() {
	var size int
	fmt.Println("enter the size of the array")
	fmt.Scanln(&size)

	myarray := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("enter the element %d", i+1)
		fmt.Scanln(&myarray[i])
	}
	fmt.Println("array elements are ", myarray)
}
