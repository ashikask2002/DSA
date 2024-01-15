package main

import "fmt"

func main() {
	arr := []int{2, 5, 4, 1, 7, 3}
	fmt.Println("array before sorting", arr)
	start := 0
	n := len(arr) - 1
	end := n

	result := QuickSort(arr, start, end)
	fmt.Println("array after sorting",result)

}

func QuickSort(arr []int,start,end int)[]int{
	if start < end {
		loc := partition(arr,start,end)

		QuickSort(arr,start,loc-1)
		QuickSort(arr,loc+1,end)
	}
	return arr
}

func partition(arr []int, start int,end int) int{
	pivot := arr[start]
	i := start+1
	j := end
	for i <j{
		for arr[i]<=pivot && i <end{
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < j{
			arr[i],arr[j]= arr[j],arr[i]
		}
	}
	arr[start],arr[j] = arr[j],arr[start]
	return j
}