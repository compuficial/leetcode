package main

import "fmt"

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.q = append(s.q, x)
	for i := 0; i < len(s.q)-1; i++ {
		s.q = append(s.q, s.q[0])
		s.q = s.q[1:]
	}
}

func (s *MyStack) Pop() int {
	val := s.q[0]
	s.q = s.q[1:]
	return val
}

func (s *MyStack) Top() int {
	return s.q[0]
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}

func main() {
	fmt.Println("Test Case 1:")
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Printf("Top: %d (Expected: 2)\n", obj.Top())
	fmt.Printf("Pop: %d (Expected: 2)\n", obj.Pop())
	fmt.Printf("Empty: %t (Expected: false)\n", obj.Empty())
	fmt.Printf("Pop: %d (Expected: 1)\n", obj.Pop())
	fmt.Printf("Empty: %t (Expected: true)\n", obj.Empty())
	fmt.Println("---")

	fmt.Println("Test Case 2: Push and Empty")
	obj2 := Constructor()
	obj2.Push(5)
	fmt.Printf("Empty: %t (Expected: false)\n", obj2.Empty())
	fmt.Println("---")

	fmt.Println("Test Case 3: Empty stack Pop/Top (will likely panic if not handled in implementation)")
	obj3 := Constructor()
	fmt.Printf("Empty: %t (Expected: true)\n", obj3.Empty())
	fmt.Println("---")
}
