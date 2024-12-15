package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)

	for i, j := 0, len(xStr)-1; i < j; i, j = i+1, j-1 {
		if xStr[i] != xStr[j] {
			return false
		}
	}

	return true
}

func main() {
	result := isPalindrome(121)
	fmt.Printf("Number: %d, isPalindrome: %v\n", 121, result)
}
