package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name, address string
	var contact = make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Input name: ")
	name, _ = reader.ReadString('\n')
	fmt.Printf("Name: %v", name)
	contact["name"] = strings.TrimSpace(name)

	fmt.Printf("Input address: ")
	address, _ = reader.ReadString('\n')
	fmt.Printf("Address: %v", address)
	contact["address"] = strings.TrimSpace(address)

	data, _ := json.Marshal(contact)
	fmt.Println(string(data))

}
