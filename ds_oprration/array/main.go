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
	insertintOarray(arr,size,2,9)
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

func  insertintOarray(arr []int, size int, index int, value int){
	for j := size; j > index; j--{
		arr[j] = arr[j-1]
	}
	arr[index] = value

}