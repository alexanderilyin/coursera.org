package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	acceleration := askNumericInput("Please enter acceleration (m/s^2): ")
	velocity := askNumericInput("Please enter velocity(m/s): ")
	initial_displacement := askNumericInput("Please enter initial displacement (m): ")

	displacementFunc := GenDisplaceFn(acceleration, velocity, initial_displacement)

	time := askNumericInput("Please enter an amount of time (s): ")
	displacement := displacementFunc(time)

	fmt.Printf("The displacement after time %fs is: %fm", time, displacement)
}

func GenDisplaceFn(acc, v_0, s_0 float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acc*math.Pow(time, 2) + v_0*time + s_0
	}
}

func askNumericInput(message string) float64 {
	var input string
	var output float64

	for {
		fmt.Print(message)
		fmt.Scanln(&input)

		number, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Printf("- INPUT: %s is not an number\n", input)
		} else {
			output = number
			break
		}
	}

	return output
}
