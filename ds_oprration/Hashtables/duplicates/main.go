package main

import "fmt"

func main() {
	arr := []int{3, 1, 5, 5, 6, 3, 7, 1}
	fmt.Println("array before ", arr)
	g := removeduplicates(arr)
	fmt.Println("array after",g)

}

func removeduplicates(arr []int) []int{
	table := make(map[int]bool)
	var array []int

	for _,v := range arr{
		if !table[v]{
			array = append(array,v)
			table[v] =true
		}

	}
	return array
}