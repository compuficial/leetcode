package main

import (
	"fmt"
	"time"
)

func twoSum(nums []int, target int) []int {
	numToIndexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, found := numToIndexMap[complement]; found {
			return []int{i, idx}
		}

		numToIndexMap[num] = i
	}

	return nil
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
