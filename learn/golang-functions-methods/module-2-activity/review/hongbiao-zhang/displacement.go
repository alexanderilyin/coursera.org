package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		s := 0.5*a*t*t + v0*t + s0
		return s
	}
	return fn
}

func main() {
	var a, v0, s0, t float64
	fmt.Printf("Enter 3 parameters seperated by whitespace: ")
	fmt.Scan(&a, &v0, &s0)
	fn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("Enter a time: ")
	fmt.Scan(&t)

	s := fn(t)
	fmt.Printf("The displacement is: %f\n", s)
}
