package main

import (
	"fmt"
)

// Define an interface type called Animal which describes the methods of an animal.
// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
// which take no arguments and return no values.

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Define three types Cow, Bird, and Snake. For each of these three types, define methods
// Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy
// the Animal interface.

type Beast struct {
	food       string
	locomotion string
	noise      string
}

// The Eat() method should print the animal’s food, the Move() method should print
// the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.

func (b Beast) Move() {
	fmt.Println(b.locomotion)
}

func (b Beast) Eat() {
	fmt.Println(b.food)
}

func (b Beast) Speak() {
	fmt.Println(b.noise)
}

type Cow struct {
	Beast
}

type Bird struct {
	Beast
}

type Snake struct {
	Beast
}

// When the user creates an animal, create an object of the appropriate type.
// Your program should call the appropriate method when the user issues a query command.

const (
	newanimal = "newanimal"
	query     = "query"
	cow       = "cow"
	bird      = "bird"
	snake     = "snake"
)

func makeAnimal(animalType string) Animal {
	switch animalType {
	case cow:
		return Cow{Beast{"grass", "walk", "moo"}}
	case bird:
		return Bird{Beast{"worms", "fly", "peep"}}
	case snake:
		return Snake{Beast{"mice", "slither", "hsss"}}
	default:
		return nil
	}
}

func main() {

	fmt.Println("Usage: 'newanimal $NAME $TYPE' or 'query $NAME $TRAIT'")

	// | Animal | Food eaten | Locomotion method | Spoken sound |
	// |--------|------------|-------------------|--------------|
	// | cow    | grass      | walk              | moo          |
	// | bird   | worms      | fly               | peep         |
	// | snake  | mice       | slither           | hsss         |

	// Your program should present the user with a prompt, “>”,
	// to indicate that the user can type a request.
	//
	// Your program should accept one command at a time from the user,
	// print out a response, and print out a new prompt on a new line.
	// Your program should continue in this loop forever.
	//
	// Every command from the user must be either a “newanimal” command
	// or a “query” command.

	var command, name, tmp string
	farm := map[string]Animal{}

	for {
		fmt.Print("> ")
		fmt.Scan(&command, &name, &tmp)

		switch command {
		case newanimal:
			// Each “newanimal” command must be a single line containing three strings.
			// The first string is “newanimal”. The second string is an arbitrary string
			// which will be the name of the new animal. The third string is the type
			// of the new animal, either “cow”, “bird”, or “snake”.
			//
			// Your program should process each newanimal command by creating the new animal
			// and printing “Created it!” on the screen.
			animal := makeAnimal(tmp)
			if animal == nil {
				fmt.Println("ERROR: Unknown animal type:", tmp)
			} else {
				farm[name] = animal
				fmt.Println("Created it!")
			}
		case query:
			// Each “query” command must be a single line containing 3 strings.
			// The first string is “query”. The second string is the name of the animal.
			// The third string is the name of the information requested about the animal,
			// either “eat”, “move”, or “speak”. Your program should process each query
			// command by printing out the requested data.
			trait := tmp
			animal, found := farm[name]
			if !found {
				fmt.Println("ERROR Can't find animal by name:", name)
			} else {
				switch trait {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("ERROR: Unknown trait:", trait)
				}

			}
		default:
			fmt.Println("ERROR: Unknown command", command)
		}
	}

}
