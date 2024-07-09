package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input    []int
		count    int
		expected [][]int
	}{
		{[]int{1, 2, 3, 4}, 4, [][]int{{1}, {2}, {3}, {4}}},
		{[]int{1, 2, 3, 4, 5}, 4, [][]int{{1, 5}, {2}, {3}, {4}}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v/%v=%v", test.input, test.count, test.expected), func(t *testing.T) {
			result := slice(test.input, test.count)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("slice(%v, %v) == %v; expected %v", test.input, test.count, result, test.expected)
			}
		})
	}
}
