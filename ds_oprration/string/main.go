package main

import (
	"fmt"
	"unicode"
)

func main() {
	word := "malayalam"

	length(word)
	upper(word)
	Palindrome(word)
}

func length(word string) {
	len := 0
	arr := []byte(word)

	for range arr {
		len++
	}
	fmt.Println("length is ", len)
}

func upper(word string) {
	var ch []rune

	for _, val := range word {
		fmt.Println("%T", val)
		ch = append(ch, unicode.ToUpper(val))
	}
	fmt.Println(string(ch))
}

func Palindrome(word string) {
	flag := 0
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			flag = 1
		}

	}
	if flag == 0 {
		fmt.Println("it is palindrome")
	} else {
		fmt.Println("it is not")
	}
}
