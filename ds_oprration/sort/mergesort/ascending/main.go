package main

import "fmt"

func main() {
	arr := []int{3, 2, 6, 5, 1, 8}
	fmt.Println("array before sorting",arr)
	start := 0
	end := len(arr) - 1
	result := MergeSort(arr,start,end)
	fmt.Println("array after sorting", result)
}


func MergeSort(arr []int,start,end int)[]int{
	
	if start < end{
		mid := (start+end)/2
		MergeSort(arr,start,mid)
		MergeSort(arr,mid+1,end)
		Merge(arr,start,mid,end)
	}
	return arr
}

func Merge(arr []int,start int,mid int,end int){
	i := start
	j := mid+1

	var b []int
	for i <= mid && j <= end{
		if arr[i] <= arr[j]{
			b  = append(b,arr[i])
			i++	
		}else {
			b = append(b,arr[j])
			j++
		}
	} 

	for ; i<= mid ; i++{
		b = append(b,arr[i])
	}
	for ; j <= end; j++{
		b = append(b,arr[j])
	}

	for i := start ; i <=  end ;i++{
      arr[i] = b[i -start]
	}
}