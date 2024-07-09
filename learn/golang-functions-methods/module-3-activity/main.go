package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

const (
	Eat   = "eat"
	Move  = "move"
	Speak = "speak"
)

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

var animals = map[string]Animal{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

func explain(pet Animal, action string) string {
	var answer string
	switch action {
	case Eat:
		answer = pet.Eat()
	case Move:
		answer = pet.Move()
	case Speak:
		answer = pet.Speak()
	default:
		answer = fmt.Sprintf("ERROR: Unknown action: %v.", action)
	}
	return answer
}

func main() {
	fmt.Println("Usage: cow|bird|snake eat|move|speak")

	for {
		var animal, action string
		fmt.Print("> ")
		fmt.Scan(&animal, &action)
		// Check if the animal exists
		pet, exists := animals[animal]
		if !exists {
			fmt.Printf("ERROR: Unknown animal: %v.\n", animal)
			continue
		}
		fmt.Println(explain(pet, action))
	}
}
