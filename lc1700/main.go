package main

import (
	"fmt"
)

func countStudents(students []int, sandwiches []int) int {
	rejected := 0
	for rejected != len(students) {
		if students[0] == sandwiches[0] {
			// fmt.Println("Student likes the sandwich, removing from both lists.")
			students = students[1:]
			sandwiches = sandwiches[1:]
			// fmt.Println("Updated students:", students)
			// fmt.Println("Updated sandwiches:", sandwiches)
			rejected = 0
		} else {
			// fmt.Println("Student does not like the sandwich, moving to the end of the line.")
			firstStudent := students[0]
			restStudents := students[1:]
			students = append(restStudents, firstStudent)
			// fmt.Println("Updated students:", students)
			// fmt.Println("Updated sandwiches:", sandwiches)
			rejected++
		}

		if rejected == len(students) {
			break
		}
	}

	return len(students)
}

func main() {
	testCases := []struct {
		students   []int
		sandwiches []int
		expected   int
	}{
		{
			students:   []int{1, 1, 0, 0},
			sandwiches: []int{0, 1, 0, 1},
			expected:   0,
		},
		{
			students:   []int{1, 1, 1, 0, 0, 1},
			sandwiches: []int{1, 0, 0, 0, 1, 1},
			expected:   3,
		},
	}

	for _, tc := range testCases {
		result := countStudents(tc.students, tc.sandwiches)
		fmt.Printf("Students: %v, Sandwiches: %v, Expected: %d, Got: %d\n", tc.students, tc.sandwiches, tc.expected, result)
		if result != tc.expected {
			fmt.Println("Test Failed!")
		} else {
			fmt.Println("Test Passed!")
		}
		fmt.Println("---")
	}
}
