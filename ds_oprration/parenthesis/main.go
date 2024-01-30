package main

import "fmt"

func main() {
    s := "{)"
    var stack = []rune{}

    remove := func(val rune) bool {
        if len(stack) == 0 {
            return false
        }

        switch {
        case val == '}':
            return stack[len(stack)-1] == '{'
        case val == ']':
            return stack[len(stack)-1] == '['
        case val == ')':
            return stack[len(stack)-1] == '('
        }

        return false
    }

   
    bracketsSatisfied := true

    for _, val := range s {
        if val == '[' || val == '(' || val == '{' {
            stack = append(stack, val)
           
        } else {
            
            if !remove(val) {
                bracketsSatisfied = false
                break
            }
            stack = stack[:len(stack)-1]
        }
    }

    if bracketsSatisfied {
        fmt.Println("Brackets are satisfied")
    } else {
        fmt.Println("Brackets are not satisfied")
    }
}
