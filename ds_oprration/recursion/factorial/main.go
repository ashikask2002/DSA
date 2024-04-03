package main

import "fmt"

func main() {
	value := 5
	result := Recursion(value)
	fmt.Println("factorial is ", result)
}

func Recursion(value int) int {
	if value <= 1 {
		return 1
	}

	return value * Recursion(value-1)
}
