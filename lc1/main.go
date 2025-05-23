package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		need := target - num

		if value, exists := seen[need]; exists {
			// fmt.Printf("Found: %d + %d = %d\n", nums[value], nums[i], target)
			// fmt.Printf("Indices: %d, %d\n", value, i)
			return []int{value, i}
		}

		seen[nums[i]] = i
	}

	return []int{}
}

func runTestCase(nums []int, target int, expected []int) {
	result := twoSum(nums, target)
	pass := reflect.DeepEqual(result, expected)
	fmt.Printf("nums: %v, target: %d => result: %v, expected: %v, pass: %v\n", nums, target, result, expected, pass)
}

func main() {
	fmt.Println("Running Two Sum Test Cases...")

	runTestCase([]int{2, 7, 11, 15}, 9, []int{0, 1})
	runTestCase([]int{3, 2, 4}, 6, []int{1, 2})
	runTestCase([]int{3, 3}, 6, []int{0, 1})
	runTestCase([]int{1, 2, 3}, 7, []int{}) // no valid pair
}
