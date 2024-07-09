package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Your program will define a name struct which has two fields,
// fname for the first name, and lname for the last name.
// Each field will be a string of size 20 (characters).

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	var names []Name

	fmt.Println("Filename to read:")
	fmt.Scanln(&filename)
	fmt.Printf("Filename: '%v'\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 2 {
			fname := parts[0]
			lname := parts[1]

			name := Name{
				fname: fname,
				lname: lname,
			}
			names = append(names, name)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.fname, name.lname)
	}

}
