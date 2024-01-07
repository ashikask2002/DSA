// // add the elements to array
// package main

// import "fmt"

// func main() {
// 	arr := make([]int, 20)
// 	fmt.Println("enter the values to array")
// 	size := 5
// 	addarray(arr, size)
// 	displayarray(arr, size)
// }
// func addarray(arr []int, size int) {
// 	for i := 0; i < size; i++ {
// 		fmt.Scanf("%d", &arr[i])
// 	}
// }

// func displayarray(arr []int, size int) {
// 	fmt.Println("the array is ")
// 	for i := 0; i < size; i++ {
// 		fmt.Printf("%d", arr[i])
// 	}
// }

package main

import "fmt"

func main() {
	arr := make([]int, 20)
	fmt.Println("enter the elements of array")
	size := 5

	Addarray(arr, size)
	// insertintOarray(arr,size,2,9)
	// Displayarray(arr, size)

	// append(arr, size, 11)

	// deleteindex(arr,size,3)
	ReverseArray(arr,size)

	Displayarray(arr, size)

}

// --------add the elements into the array-------

func Addarray(arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
}

func Displayarray(arr []int, size int) {
	fmt.Println("the entered elements are ")
	for i := 0; i < size; i++ {
		fmt.Printf("%d\n", arr[i])
	}
}

// ----- insert into an array-------

func insertintOarray(arr []int, size int, index int, value int) {
	for j := size; j > index; j-- {
		arr[j] = arr[j-1]
	}
	arr[index] = value

}

// ------append-------//
func append(arr []int, size int, value int) {
	arr[size-1] = value
	fmt.Println(value)

}

// --------deleteindex-------
func deleteindex(arr []int, size int, index int) {
	for i := 0; i < size; i++ {
		if i == index {
			for j := i; j < size; j++ {
				arr[j] = arr[j+1]
			}
			size = size - 1
		}
	}
}

func ReverseArray(arr[]int,size int){
	var temp int
	for i,j := 0,size-1; i<j;i,j = i+1,j-1{
		temp = arr[i]
		arr[i]=arr[j]
		arr[j]=temp
	}
}