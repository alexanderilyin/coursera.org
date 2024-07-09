package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (animal Cow) Eat()   { fmt.Println("grass") }
func (animal Cow) Move()  { fmt.Println("walk") }
func (animal Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (animal Bird) Eat()   { fmt.Println("worms") }
func (animal Bird) Move()  { fmt.Println("fly") }
func (animal Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (animal Snake) Eat()   { fmt.Println("mice") }
func (animal Snake) Move()  { fmt.Println("slither") }
func (animal Snake) Speak() { fmt.Println("hsss") }

type Farm map[string]Animal

func (farm Farm) CreateAnimal(animalType string, name string) {
	switch animalType {
	case "cow":
		farm[name] = Cow{}
	case "bird":
		farm[name] = Bird{}
	case "snake":
		farm[name] = Snake{}
	default:
		fmt.Println("unknown animal type")
	}
	fmt.Println("Created it!")
}

func (farm Farm) QueryAnimal(name string, action string) {
	animal, present := farm[name]

	if present {
		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("unknown action")
		}
	} else {
		fmt.Println("animal does not exist")
	}
}

func main() {
	myFarm := Farm{}

	fmt.Println("operation 1: > newanimal <name> <anmial>")
	fmt.Println("operation 2: > query <name> <action>")
	for true {
		command := ReadCommand()
		ProcessCommand(myFarm, command)
	}
}

func ProcessCommand(farm Farm, command map[string]string) {
	switch command["operation"] {
	case "newanimal":
		farm.CreateAnimal(command["animal"], command["name"])
	case "query":
		farm.QueryAnimal(command["name"], command["action"])
	default:
		fmt.Println("unknown operation")
	}
}

func ReadCommand() map[string]string {
	var operation, name, animal, action, errorMessage string

	line := ReadLine()
	operation, name, animal, action, errorMessage = ValidateLine(line)
	for errorMessage != "" {
		fmt.Println(errorMessage)
		fmt.Println("operation 1: > newanimal <name> <anmial>")
		fmt.Println("operation 2: > query <name> <action>")
		line := ReadLine()
		operation, name, animal, action, errorMessage = ValidateLine(line)
	}

	return map[string]string{"operation": operation, "name": name, "animal": animal, "action": action}
}

func ReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	userInput := ""

	fmt.Printf("> ")
	if scanner.Scan() {
		userInput = scanner.Text()
	}

	return userInput
}

func ValidateLine(line string) (string, string, string, string, string) {
	validOperations := []string{"newanimal", "query"}
	validAnimals := []string{"cow", "bird", "snake"}
	validActions := []string{"eat", "move", "speak"}

	words := strings.Fields(line)
	if len(words) != 3 {
		return "", "", "", "", "you should enter three words."
	}
	if !InArray(words[0], validOperations) {
		return "", "", "", "", "not a valid operation, must be one of these: " + strings.Join(validOperations, ", ")
	}
	if words[0] == "newanimal" {
		if !InArray(words[2], validAnimals) {
			return "", "", "", "", "not a valid animal, must be one of these: " + strings.Join(validAnimals, ", ")
		}
		return words[0], words[1], words[2], "", ""
	}
	if words[0] == "query" {
		if !InArray(words[2], validActions) {
			return "", "", "", "", "not a valid actions, must be one of these: " + strings.Join(validActions, ", ")
		}
		return words[0], words[1], "", words[2], ""
	}

	return "", "", "", "", "something went wrong"
}

func InArray(data string, dataArray []string) bool {
	for _, v := range dataArray {
		if v == data {
			return true
		}
	}

	return false
}
