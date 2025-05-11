package main

import "fmt"

func main() {
	fmt.Println("Static Arrays")
	fmt.Println("-------------")
	fmt.Println("The Array")
	iA := [6]int{3, 5, 7, 1, 6}
	fmt.Println(iA)

	fmt.Println("inserted 4 at index 2")
	// inserting at the end
	iA[len(iA)-1] = 4
	fmt.Println(iA)

	// reading from index
	fmt.Println("Reading from index 2")
	fmt.Println(iA[2])

	// traversing the array
	fmt.Println("Looping through array")
	for i := 0; i < len(iA); i++ {
		fmt.Println(iA[i])
	}

	fmt.Println("Removing from the end of the array")
	iA[len(iA)-1] = 0
	fmt.Println(iA)

	fmt.Println("Insert 9 in the middle of the array at index 2")
	insertMiddle(iA[:], 2, 9, len(iA)-1)
	fmt.Println(iA)
}

func insertMiddle(arr []int, i, n, length int) []int {
	for index := length - 1; index > i-1; index-- {
		arr[index+1] = arr[index]
	}

	arr[i] = n
	return arr
}

func removeMidlde(arr []int, i, length int) []int {
	for index := i + 1; index < length; index++ {
		arr[index-1] = arr[index]
	}
	arr[length-1] = 0
	return arr
}
