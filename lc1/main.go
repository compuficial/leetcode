package main

import (
	"fmt"
	"slices"
	"time"
)

func twoSum(nums []int, target int) []int {
	var result []int
	for idx, num := range nums {
		num1 := target - num
		num1 = slices.Index(nums, num1)
		if num1 == -1 || num1 == idx {
			continue
		}

		return []int{idx, num1}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	start := time.Now()
	result := twoSum(nums, target)
	elapsed := time.Since(start)
	fmt.Println("time:", elapsed)
	fmt.Println("result:", result)
}
