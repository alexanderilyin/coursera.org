package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	println("This animal eats", animal.food)
}

func (animal Animal) Move() {
	println("This animal locomotion method is", animal.locomotion)
}

func (animal Animal) Speak() {
	println("This animal spoken sound is", animal.noise)
}

func main() {

	var inputName, inputInformation string
	var animals = make(map[string]Animal)
	animals["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for true {

		fmt.Printf(">")
		fmt.Scan(&inputName, &inputInformation)
		switch strings.ToLower(inputInformation) {
		case "eat":
			animals[inputName].Eat()
		case "move":
			animals[inputName].Move()
		case "speak":
			animals[inputName].Speak()
		default:
			fmt.Println("This option is not supported.")
		}

	}
}
