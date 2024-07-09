package main

import (
	"testing"
)

func TestGenDisplaceFn(t *testing.T) {
	tests := []struct {
		acceleration float64
		velocity     float64
		displacement float64
		time         float64
		expected     float64
	}{
		{10, 2, 1, 0, 1},
		{10, 2, 1, 3, 52},
		{10, 2, 1, 5, 136},
	}

	for _, test := range tests {
		fn := GenDisplaceFn(test.acceleration, test.velocity, test.displacement)
		actual := fn(test.time)
		if actual != test.expected {
			t.Errorf("GenDisplaceFn(%v, %v, %v)(%v) = %v; expected %v",
				test.acceleration,
				test.velocity,
				test.displacement,
				test.time,
				actual,
				test.expected)
		}
	}
}
