package main

import (
	"fmt"
	"strings"
)

// Write a program which prompts the user to enter a string.
// The program searches through the entered string for the characters
// ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered
// string starts with the character ‘i’, ends with the character ‘n’,
// and contains the character ‘a’. The program should print “Not Found!”
// otherwise. The program should not be case-sensitive, so it does not
// matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example
// entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
// The program should print “Not Found!” for the following strings,
// “ihhhhhn”, “ina”, “xian”.

func main() {
	var s string
	fmt.Printf("Input string: ")
	fmt.Scan(&s)
	fmt.Printf("Input Type: '%T', Input Value '%v'\n", s, s)

	s = strings.ToLower(s)

	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
