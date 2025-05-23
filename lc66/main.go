package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func plusOne_S2(digits []int) []int {
	i := len(digits) - 1
	digits[i]++

	for digits[i] == 10 {
		digits[i] = 0
		i--

		if i < 0 {
			return append([]int{1}, digits...)
		}
		digits[i]++
	}

	return digits
}

func plusOne_S1(digits []int) []int {
	digits[len(digits)-1]++
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Printf("i:%d, num: %d\n", i, digits[i])
		if carry == 1 {
			digits[i]++
		}

		if digits[i] == 10 {
			carry = 1
			digits[i] = 0
		} else {
			carry = 0
		}
	}

	if carry == 1 {
		newDigits := make([]int, len(digits)+1)
		copy(newDigits[1:], digits)
		newDigits[0] = 1
		return newDigits
	}

	return digits
}

func main() {
	// Test Case 1
	digits1 := []int{1, 2, 3}
	fmt.Printf("Input: %v\n", digits1)
	result1 := plusOne(digits1)
	fmt.Printf("Output: %v\n\n", result1) // Expected: [1, 2, 4]

	// Test Case 2
	digits2 := []int{4, 3, 2, 1}
	fmt.Printf("Input: %v\n", digits2)
	result2 := plusOne(digits2)
	fmt.Printf("Output: %v\n\n", result2) // Expected: [4, 3, 2, 2]

	// Test Case 3
	digits3 := []int{9}
	fmt.Printf("Input: %v\n", digits3)
	result3 := plusOne(digits3)
	fmt.Printf("Output: %v\n\n", result3) // Expected: [1, 0]

	// Test Case 4
	digits4 := []int{1, 9, 9}
	fmt.Printf("Input: %v\n", digits4)
	result4 := plusOne(digits4)
	fmt.Printf("Output: %v\n\n", result4) // Expected: [2, 0, 0]

	// Test Case 5
	digits5 := []int{9, 9, 9, 9}
	fmt.Printf("Input: %v\n", digits5)
	result5 := plusOne(digits5)
	fmt.Printf("Output: %v\n\n", result5) // Expected: [1, 0, 0, 0, 0]
}
