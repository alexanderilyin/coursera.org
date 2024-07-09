package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	// s = ½*a*t^2 + vo*t + so
	fn := func(t float64) float64 {
		return a*math.Pow(t, 2)/2 + vo*t + so
	}
	return fn
}

func main() {
	var a, vo, so float64

	fmt.Print("Acceleration (a): ")
	fmt.Scan(&a)

	fmt.Print("Initial Velocity (vo): ")
	fmt.Scan(&vo)

	fmt.Print("Initial Displacement (so): ")
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, vo, so) // 10, 2, 1

	var t float64

	fmt.Print("Time (t): ")
	fmt.Scan(&t)

	fmt.Println("Displacement (s): ", fn(t)) // 3
	// fmt.Println("Displacement: ", fn(5))
}
