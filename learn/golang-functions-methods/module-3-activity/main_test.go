package main

import (
	"testing"
)

func TestAnimals(t *testing.T) {
	tests := []struct {
		animal   string
		action   string
		expected string
	}{
		{"cow", "eat", "grass"},
		{"bird", "move", "fly"},
		{"snake", "speak", "hsss"},
	}

	for _, test := range tests {
		actual := explain(animals[test.animal], test.action)
		if actual != test.expected {
			t.Errorf("%v %v:%v; expected %v",
				test.animal,
				test.action,
				actual,
				test.expected,
			)
		}
	}
}
