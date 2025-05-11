package main

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	an := 2 * n
	aNums := make([]int, an)

	for i := 0; i < len(nums); i++ {
		aNums[i] = nums[i]
		aNums[i+n] = nums[i]
	}

	return aNums
}

func main() {
	fmt.Println("LC1929: Concatenation of Array")
	fmt.Println("------------------------------")

	nums := []int{1, 2, 1}
	aNums := getConcatenation(nums)
	fmt.Printf("nums: %v, aNums %v\n", nums, aNums)
}
