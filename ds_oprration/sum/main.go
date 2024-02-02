package main

import "fmt"

func findpair(arr []int, targetsun int) (int, int, bool) {
	for i := 0; i <= len(arr)-1; i++ {
		for j := i + 1; i <= len(arr); j++ {
			if arr[i]+arr[j] == targetsun {
				return arr[i], arr[j], true
			}
		}
	}
	return 0, 0, false
}

func main() {
	myArray := []int{2, 5, 4, 8, 1}
	targersum := 10

	num1, num2, found := findpair(myArray, targersum)
	if found {
		fmt.Printf("pair we get is %d, and %d", num1, num2)
	} else {
		fmt.Println("didnt get the pairs")
	}
}
