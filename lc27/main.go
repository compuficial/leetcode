package main

import "fmt"

func main() {
	fmt.Println("LC27: Remove Element")
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums, k := removeElement(nums, 2)
	fmt.Println("array:", nums)
	fmt.Println("k:", k)
}

func removeElement(nums []int, val int) ([]int, int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return nums, k
}
