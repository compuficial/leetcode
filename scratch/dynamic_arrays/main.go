package main

import "fmt"

type DynamicArray struct {
	Capacity, Length int
	Arr              []int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		Capacity: 4,
		Length:   0,
		Arr:      make([]int, 4),
	}
}

func (d *DynamicArray) Pushback(n int) {
	if d.Length == d.Capacity {
		d.Resize()
	}

	d.Arr[d.Length] = n
	d.Length++
}

func (d *DynamicArray) Resize() {
	d.Capacity *= 2
	newArr := make([]int, d.Capacity)

	copy(newArr, d.Arr)
	d.Arr = newArr
}

func (d *DynamicArray) Popback() {
	if d.Length > 0 {
		d.Length--
		d.Arr[d.Length] = 0 // zero out popped value
	}
}

func (d *DynamicArray) Get(i int) int {
	if i >= 0 && i < d.Length {
		return d.Arr[i]
	}

	return -1
}

func (d *DynamicArray) Insert(i, n int) {
	if i < 0 || i > d.Length {
		return
	}

	if d.Length == d.Capacity {
		d.Resize()
	}

	for j := d.Length; j > i; j-- {
		d.Arr[j] = d.Arr[j-1]
	}

	d.Arr[i] = n
	d.Length++
}

func (d *DynamicArray) Print() {
	for i := 0; i < d.Length; i++ {
		fmt.Println(d.Arr[i])
	}
}

func main() {
	fmt.Println("Dynamic Arrays")
	fmt.Println("-------------")
	arr := NewDynamicArray()

	fmt.Println("arr:", arr)
	arr.Pushback(2)
	arr.Pushback(1)
	fmt.Println("arr:", arr)
	arr.Insert(0, 2)
	arr.Insert(1, 3)
	fmt.Println("arr:", arr)
	fmt.Println("arr:", arr)
	arr.Popback()
	arr.Popback()
	fmt.Println("arr:", arr)
}
