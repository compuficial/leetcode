package main

import "fmt"

func main() {
	fmt.Println("LC26: Remove Duplicates from Sorted Array")
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums, k := removeDuplicates(nums)
	fmt.Println("array:", nums)
	fmt.Println("k:", k)
}

func removeDuplicates(nums []int) ([]int, int) {
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return nums, k
}
