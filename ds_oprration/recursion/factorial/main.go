package main

import "fmt"

func main() {
	value := 6
	result := Recursion(value)
	fmt.Printf("value is %d", result)

}

func Recursion(value int) int {
	if value == 0 {
		return 1
	}
	return value * Recursion(value-1)
}
