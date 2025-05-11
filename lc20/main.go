package main

import "fmt"

// ({[]})
func isValid(s string) bool {
	var stack []rune

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else if r == ')' || r == ']' || r == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (r == ')' && top != '(') ||
				(r == ']' && top != '[') ||
				(r == '}' && top != '{') {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println("LC20: Valid Parentheses")
	fmt.Println("-----------------------")
	v := isValid("([{}])")
	fmt.Println("isValid:", v)
}
