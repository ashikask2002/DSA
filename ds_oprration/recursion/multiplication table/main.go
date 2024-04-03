package main

import "fmt"

func main() {

	number := 5
	limit := 10

	multipication(number,limit,1)
}

func multipication(number, limit, current int) {
	if limit < current {
		return
	}

	value := number * current
	fmt.Printf("%d * %d = %d \n", number, current, value)

	multipication(number,limit,current+1)
}
