package main

import (
	"fmt"
)

func searchInsertSearch(nums []int, target int) int {
	idx := 0
	for _, num := range nums {
		if num >= target {
			return idx
		}

		idx++
	}

	return idx
}

func searchInsertS2(nums []int, target int) int {
	idx := 0
	for i, num := range nums {
		if num == target {
			return i
		}

		if num <= target {
			idx++
		}
	}

	return idx
}

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}

		if num < target && i == len(nums)-1 {
			return i + 1
		} else if num < target && target < nums[i+1] {
			return i + 1
		}
	}
	return 0
}

func runTestCase(nums []int, target int, expected int) {
	result := searchInsert(nums, target)
	pass := result == expected
	fmt.Printf("nums: %v, target: %d => result: %d, expected: %d, pass: %v\n", nums, target, result, expected, pass)
}

func main() {
	fmt.Println("Running Search Insert Position Test Cases...")

	runTestCase([]int{1, 3, 5, 6}, 5, 2)
	runTestCase([]int{1, 3, 5, 6}, 2, 1)
	runTestCase([]int{1, 3, 5, 6}, 7, 4)
	runTestCase([]int{1, 3, 5, 6}, 0, 0)
	runTestCase([]int{1}, 0, 0)
}
