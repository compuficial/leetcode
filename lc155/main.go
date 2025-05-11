package main

import "fmt"

type MinStack struct {
	Capacity, Length int
	Arr, Min         []int
}

func Constructor() *MinStack {
	return &MinStack{
		Capacity: 4,
		Length:   0,
		Arr:      make([]int, 4),
		Min:      make([]int, 4),
	}
}

func (ms *MinStack) Resize() {
	ms.Capacity *= 2
	newArr := make([]int, ms.Capacity)
	copy(newArr, ms.Arr)
	ms.Arr = newArr

	newMin := make([]int, ms.Capacity)
	copy(newMin, ms.Min)
	ms.Min = newMin
}

func (ms *MinStack) Push(val int) {
	if ms.Length == ms.Capacity {
		fmt.Println("Resizing...")
		ms.Resize()
	}

	ms.Arr[ms.Length] = val

	if ms.Length == 0 {
		ms.Min[ms.Length] = val
	} else {
		currMin := ms.Min[ms.Length-1]

		if val < currMin {
			ms.Min[ms.Length] = val
		} else {
			ms.Min[ms.Length] = currMin
		}
	}

	ms.Length++
}

func (ms *MinStack) Pop() {
	if ms.Length == 0 {
		return
	}

	ms.Length--
	ms.Arr[ms.Length] = 0
	ms.Min[ms.Length] = 0
}

func (ms *MinStack) Top() int {
	return ms.Arr[ms.Length-1]
}

func (ms *MinStack) GetMin() int {
	return ms.Min[ms.Length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	fmt.Println("LC155: Min Stack")
	fmt.Println("----------------")

	ms := Constructor()
	ms.Push(5)
	fmt.Println(ms)
	ms.Push(3)
	fmt.Println(ms)
	ms.Push(9)
	fmt.Println(ms)
	ms.Push(4)
	fmt.Println(ms)
	fmt.Println("top:", ms.Top())
	ms.Pop()
	fmt.Println("pop:", ms)
	fmt.Println("top", ms.Top())
	fmt.Println(ms)
	ms.Push(6)
	fmt.Println(ms)
	ms.Push(7)
	fmt.Println(ms)
}
