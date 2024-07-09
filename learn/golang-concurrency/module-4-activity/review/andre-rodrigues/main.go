package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	nbr     int
	leftCS  *ChopS
	rightCS *ChopS
}

type Host struct {
	name string
}

func (p Philo) eat() {
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Printf("starting to eat %d\n", p.nbr)

	fmt.Printf("Eating %d\n", p.nbr)
	fmt.Printf("Eating %d\n", p.nbr)
	fmt.Printf("Eating %d\n", p.nbr)

	fmt.Printf("Finished eating %d\n", p.nbr)
	p.leftCS.Unlock()
	p.rightCS.Unlock()
}

func (p Philo) requestPermission(c chan bool) {
	fmt.Printf("%d - Requesting permission to host\n", p.nbr)
	c <- true
}

func (p Philo) communicateEnding(c chan bool) {
	fmt.Printf("%d - Communicating done eating\n", p.nbr)
	c <- true
}

func (h Host) eatingPlaybook(c1 chan bool, c2 chan bool, c3 chan bool, c4 chan bool, c5 chan bool) {
	/*
		This playbook could be a bit more generalized, what we want to ensure is that 2 adjacent philosophers don't eat at the same time.
	*/
	// 1 and 3 eat first -  lock CS 0,1 and 2,3
	b1 := <-c1
	b3 := <-c3
	if b1 && b3 {
		fmt.Println("Host allowed start of Philo 1 and 3")
	}
	// once 1 finishes - 5 starts - unlock 0,1, lock 4,0
	<-c1
	b5 := <-c5
	if b5 {
		fmt.Println("Philo 1 finished, Philo 5 started")
	}
	// once 3 finishes - 2 starts - unlock 2,3, lock 1,2
	<-c3
	b2 := <-c2
	if b2 {
		fmt.Println("Philo 3 finished, Philo 2 started")
	}

	// once 5 finishes - 4 starts - unlock 4,0, lock 3,4
	<-c5
	b4 := <-c4
	if b4 {
		fmt.Println("Philo 5 finished, Philo 3 started")
	}
	// close c2 and c4
	<-c4
	<-c2
}

func runHost(h Host, c1 chan bool, c2 chan bool, c3 chan bool, c4 chan bool, c5 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting host")
	h.eatingPlaybook(c1, c2, c3, c4, c5)
	fmt.Println("Ended Host")
}

func runPhilo(p Philo, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting Philo %d\n", p.nbr)
	p.requestPermission(ch)
	p.eat()
	p.communicateEnding(ch)
	fmt.Printf("Philo %d terminated\n", p.nbr)
}

func main() {
	var wg sync.WaitGroup
	permissionChan1 := make(chan bool)
	permissionChan2 := make(chan bool)
	permissionChan3 := make(chan bool)
	permissionChan4 := make(chan bool)
	permissionChan5 := make(chan bool)

	fmt.Println("Starting")
	//initialize the 5 Chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	//initialize the 5 Philos
	Philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		Philos[i] = &Philo{nbr: i + 1, leftCS: CSticks[i], rightCS: CSticks[(i+1)%5]}
	}
	host := Host{name: "host"}
	wg.Add(6)
	//In here we have 1 channel for each philosopher but this can possibly be done without the "Exclusive channels"
	go runHost(host, permissionChan1, permissionChan2, permissionChan3, permissionChan4, permissionChan5, &wg)
	go runPhilo(*Philos[0], permissionChan1, &wg)
	go runPhilo(*Philos[1], permissionChan2, &wg)
	go runPhilo(*Philos[2], permissionChan3, &wg)
	go runPhilo(*Philos[3], permissionChan4, &wg)
	go runPhilo(*Philos[4], permissionChan5, &wg)
	wg.Wait()
}
