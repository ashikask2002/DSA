package main

import "fmt"

func main() {
	str := "ashik"
	strconv := []byte(str)
	s := 0
	e := len(str) - 1

	reslt := Recursion(strconv, s, e)
	fmt.Println(string(reslt))

}

func Recursion(str []byte, s int, e int) []byte {
	if s > e {
		return str
	}
	temp := str[s]
	str[s] = str[e]
	str[e] = temp

	return Recursion(str, s+1, e-1)
}
