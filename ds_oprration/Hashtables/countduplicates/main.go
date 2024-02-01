package main

import "fmt"

func main() {
	arr := []int{4, 1, 6, 6, 2, 2, 8, 1}
	fmt.Println("before array is ", arr)
	a := removeDuplicates(arr)
	fmt.Println("count of duplicates is ",a)
}
func removeDuplicates(arr []int) int{
	table := make(map[int]int)
	count := 0

	for _,v:= range arr{
		table[v]++
		if table[v] == 2{
			count++
		}
	}
	return count
}
