package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

// Write a program to sort an array of integers. The program should partition the array into 4 parts,
// each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array
// should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

func slice(array []int, count int) (slices [][]int) {
	// slices := [][]int{a, b, c, d} // Array of slices
	slices = make([][]int, count)
	for i, val := range array {
		slices[i%count] = append(slices[i%count], val)
	}
	// Return the slices in the correct order
	return
}

func sorter(slices *[][]int, slice int, wg *sync.WaitGroup) {
	sort.Ints((*slices)[slice])
	wg.Done()
}

func main() {
	var input []int
	var symbol string

	fmt.Println("Usage: Enter list of numbers to be sorted. Provide one integer at a time. Enter X to start sorting.")

	for {
		fmt.Print("> ")
		fmt.Scan(&symbol)

		if symbol == "X" {
			break
		}

		if number, error := strconv.Atoi(symbol); error == nil {
			input = append(input, number)
		}
	}

	fmt.Println("input: ", input)

	slices := slice(input, 4)

	fmt.Println("slices: ", slices)

	var wg sync.WaitGroup

	for i := range slices {
		wg.Add(1)
		go sorter(&slices, i, &wg)
	}
	wg.Wait()

	fmt.Println("sorted slices: ", slices)

	merged := []int{}
	for _, slice := range slices {
		merged = append(merged, slice...)
	}

	sort.Ints(merged)

	fmt.Println("merged & sorted: ", merged)
}
