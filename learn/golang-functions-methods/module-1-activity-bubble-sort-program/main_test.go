package main

import (
	"reflect"
	"testing"
)

// TestSwap tests the Swap function
func TestSwap(t *testing.T) {
	tests := []struct {
		input    []int
		index    int
		expected []int
	}{
		{[]int{5, 3}, 0, []int{3, 5}},
		{[]int{1, 2, 3, 4}, 1, []int{1, 2, 3, 4}},
		{[]int{9, 8, 7, 6}, 2, []int{9, 8, 6, 7}},
	}

	for _, test := range tests {
		Swap(test.input, test.index)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Swap(%v, %d) = %v; expected %v", test.input, test.index, test.input, test.expected)
		}
	}
}

// TestBubbleSort tests the BubbleSort function
func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 6, 2}, []int{2, 3, 5, 6, 8}},
		{[]int{1, 4, 2, 3, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{9, 7, 5, 3, 1}, []int{1, 3, 5, 7, 9}},
		{[]int{9, 7, 5, 3, -1}, []int{-1, 3, 5, 7, 9}},
	}

	for _, test := range tests {
		BubbleSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("BubbleSort(%v) = %v; expected %v", test.input, test.input, test.expected)
		}
	}
}
