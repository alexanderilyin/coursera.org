package main

// Write a program which prompts the user to enter a floating point number
// and prints the integer which is a truncated version of the floating point number
// that was entered. Truncation is the process of removing the digits to the right
// of the decimal place.

import (
	"fmt"
	"strconv"
)

func main() {
	var x string

	fmt.Printf("Enter floating point number: ")
	fmt.Scan(&x)

	fmt.Printf("Input Type: %T; Input Value: %v\n", x, x)

	y, err := strconv.ParseFloat(x, 32)
	if err != nil {
		fmt.Println("Error parsing input:", err)
	}

	var z int
	z = int(y)
	fmt.Printf("Truncated: %v\n", z)
}
