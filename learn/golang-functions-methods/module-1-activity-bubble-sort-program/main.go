package main

// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements
// are in sorted order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

import (
	"fmt"
)

func BubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			Swap(numbers, j)
		}
	}

}

func Swap(numbers []int, index int) {
	if numbers[index] > numbers[index+1] {
		numbers[index], numbers[index+1] = numbers[index+1], numbers[index]
	}
}

func main() {
	var data = make([]int, 0, 10)
	var tmp int
	for i := 0; i < 10; i++ {
		fmt.Print("Input number: ")
		fmt.Scan(&tmp)
		data = append(data, tmp)
	}

	// TODO: 2. Write two tests
	fmt.Printf("Input: %v\n", data)
	BubbleSort(data)
	fmt.Printf("Sorted: %v\n", data)
}
