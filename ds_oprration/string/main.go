package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	word := "malayalam"

	length(word)
	upper(word)
	Palindrome(word)
	// removeWowels(word)
	// reverseString(word)
}

func length(word string){
	arr := []byte(word)
    length := 0
	for range arr{
		length++
	}
	fmt.Println("length is ",length)
}

func upper(word string){
	var ch []rune 

	for _,value := range word{
     ch = append(ch,unicode.ToUpper(value))
	}
	fmt.Println(string(ch))
}

func Palindrome(word string){
	flag := 0
	for i,j := 0,len(word)-1;i<j;i,j = i+1,j-1{
    if word[i]!= word[j]{
		flag = 1

	}
	}
	if flag == 0{
		fmt.Println("it is palindrome")
	}else {
		fmt.Println("it is not palindrome")
	}
}

func removeWowels(word string) {
	word2 := ""

	for _, val := range word {
		if string(val) != "a" && string(val) != "e" && string(val) != "i" && string(val) != "o" && string(val) != "u" {
			word2 += string(val)
		}
	}
	fmt.Println(string(word2))
}

func reverseString(word string) {

	words := strings.Fields(word)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	str := strings.Join(words, " ")
	fmt.Println(str)
}
