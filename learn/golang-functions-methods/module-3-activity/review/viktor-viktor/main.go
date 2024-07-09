package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a program which allows the user to get information about a predefined set of animals.
// Three animals are predefined,
//		[X] - cow,
//		[X] - bird,
//		[X] - and snake.
//	Each animal can \
//		[X] - eat,
//		[X] - move, and
//		[X] - speak.
//	The user can issue a request to find out one of three things about an animal:
//		1) the food that it eats,
//		2) its method of locomotion, and
//		3) the sound it makes when it speaks.
// The following table contains the three animals and their associated data which should be hard-coded into your program.
//
//		|Animal	|Food eaten	|Locomotion method 	|Spoken sound	|
//		|cow	|grass		|walk				|moo			|
//		|bird	|worms		|fly				|peep			|
//		|snake	|mice		|slither			|hsss			|
//
//	[X] Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
//	[X] Your program accepts one request at a time from the user,
//	[X] prints out the answer to the request, and prints out a new prompt.
//	[X] Your program should continue in this loop forever.
//	[X] Every request from the user must be a single line containing 2 strings.
//	The first string is the name of an animal,
//		- either “cow”, “bird”, or “snake”.
//	The second string is the name of the information requested about the animal,
//		- either “eat”, “move”, or “speak”.
//	[X] Your program should process each request by printing out the requested data.
//
//	You will need a data structure to hold the information about each animal.
//	Make a type called
//			- Animal
//	which is a struct containing three fields:
//		- food,
//		- locomotion, and
//		- noise,
//	all of which are strings.
//	Make three methods called
//		- Eat(),
//		- Move(), and
//		- Speak().
//	The receiver type of all of your methods should be your Animal type.
//		- The Eat() method should print the animal’s food,
//		- the Move() method should print the animal’s locomotion, and
//		- the Speak() method should print the animal’s spoken sound.
//	Your program should call the appropriate method when the user makes a request.

// Test the program by running it and testing it by issuing three requests.
// [X] - Each request should involve a different animal (cow, bird, snake) and
//		a different property of the animal (eat, move, speak).
// 			- Give 2 points for each request for which the program provides the correct response.
//
// Examine the code.
//		[X] - If the code contains a type called Animal, which is a struct containing three fields, all of which are strings, then give another 2 points.
//		- If the program contains three methods called Eat(), Move(), and Speak(), and all of their receiver types are Animal, give another 2 points.

type Animal struct {
	// - food,
	food string
	// - locomotion, and
	locomotion string
	// - noise,
	noise string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	//		|Animal	|Food eaten	|Locomotion method 	|Spoken sound	|
	//		|cow	|grass		|walk				|moo			|
	//		|bird	|worms		|fly				|peep			|
	//		|snake	|mice		|slither			|hsss			|

	farm := make(map[string]Animal)

	farm["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	farm["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	farm["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var nameOfAnimal string
		var nameOfInformationType string

		fmt.Print("type your request 'cow eat', 'bird move' etc.: >")
		if scanner.Scan() {
			request := scanner.Text()
			requestArr := strings.Split(request, " ")

			if len(requestArr) == 2 {
				nameOfAnimal = requestArr[0]
				nameOfInformationType = requestArr[1]
			} else {
				fmt.Println("bad request. try again")
				continue
			}

		}

		//fmt.Print("type name of animal[“cow”, “bird”, or “snake”]: >")
		//if scanner.Scan() {
		//	nameOfAnimal = scanner.Text()
		//}

		animal, ok := farm[nameOfAnimal]
		if ok != true {
			fmt.Println("unfortunately, there is no animal with such name on our farm: ", nameOfAnimal)
			continue
		}

		//fmt.Print("type information you want to know about ", nameOfAnimal, "[“eat”, “move”, or “speak”]: >")
		//if scanner.Scan() {
		//	nameOfInformationType = scanner.Text()
		//}

		switch nameOfInformationType {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("unfortunately, there is no such information about the animal: ", nameOfAnimal)
			continue
		}
	}

}
