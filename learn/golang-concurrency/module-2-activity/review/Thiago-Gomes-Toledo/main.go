package main

import (
	"fmt"
	"time"
)

type Element struct {
	value int
}

func main() {

	var elemOne, elemTwo Element
	elemOne.value = 1
	elemTwo.value = 2

	go setSumAtBoth(elemOne, elemTwo)

	go subtractConcurrency(elemOne, elemTwo)

	time.Sleep(2 * time.Second)

}

func setSumAtBoth(elemOne Element, elemTwo Element) {
	fmt.Println("Sum - init")
	var sum = elemOne.value + elemTwo.value

	elemOne.value = sum
	elemTwo.value = sum

	fmt.Println("Sum: ", sum)
	fmt.Println("SUM - elemOne: ", elemOne.value)
	fmt.Println("SUM - elemTwo: ", elemTwo.value)
}

func subtractConcurrency(elemOne Element, elemTwo Element) {
	fmt.Println("Subtract - init")
	var subtract = elemOne.value - elemTwo.value

	elemOne.value = subtract
	elemTwo.value = subtract

	fmt.Println("Subtract: ", subtract)
	fmt.Println("Subtract - elemOne: ", elemOne.value)
	fmt.Println("Subtract - elemTwo: ", elemTwo.value)
}

// The race conditions is when two or more goroutines access the same variable at the same time.
// In this case, the variable elemOne and elemTwo are accessed by two goroutines at the same time.
// Since both goroutines are accessing and updating elemOne and elemTwo value in a concurrent way, the
// order is non deterministic, so the output when printed can be different each time, as we can see next:

// Some Output:
// PC-tgtoledo:tarefa tgtoledo$ go run goroutines.go
// Subtract - init
// Subtract:  -1
// Subtract - elemOne:  -1
// Subtract - elemTwo:  -1
// Sum - init
// Sum:  3
// SUM - elemOne:  3
// SUM - elemTwo:  3
// PC-tgtoledo:tarefa tgtoledo$ go run goroutines.go
// Subtract - init
// Subtract:  -1
// Subtract - elemOne:  -1
// Subtract - elemTwo:  -1
// Sum - init
// Sum:  3
// SUM - elemOne:  3
// SUM - elemTwo:  3
// PC-tgtoledo:tarefa tgtoledo$ go run goroutines.go
// Sum - init
// Sum:  3
// SUM - elemOne:  3
// Subtract - init
// Subtract:  -1
// Subtract - elemOne:  -1
// Subtract - elemTwo:  -1
// SUM - elemTwo:  3
