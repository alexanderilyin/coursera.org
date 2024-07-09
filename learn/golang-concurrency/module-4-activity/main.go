package main

import (
	"fmt"
	"sync"
	"time"
)

type table struct {
	chairs    []*philosopher
	tableware []*chopstick
	host      *host
}

func (t *table) feast() {
	var wg sync.WaitGroup
	fmt.Println("Let the feeding begin!")
	for _, p := range t.chairs {
		wg.Add(1)
		go p.eat(&wg, t.host)
	}
	wg.Wait()
}

type host struct {
	limit  int
	eating chan struct{}
}

func (h *host) setLimits() {
	h.eating = make(chan struct{}, h.limit)
}

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	number              int
	meals               int
	leftHand, rightHand *chopstick
}

func (p *philosopher) askPermission(h *host) {
	h.eating <- struct{}{}
}

func (p *philosopher) eat(wg *sync.WaitGroup, h *host) {
	defer wg.Done()

	for i := 1; i <= p.meals; i++ {
		// The philosophers pick up the chopsticks in any order,
		// not lowest-numbered first (which we did in lecture).
		p.leftHand.Lock()
		p.rightHand.Lock()

		// In order to eat, a philosopher must get permission from a host
		// which executes in its own goroutine.
		go p.askPermission(h)

		// When a philosopher starts eating (after it has obtained necessary locks)
		// it prints “starting to eat <number>” on a line by itself, where <number>
		// is the number of the philosopher.
		fmt.Printf("starting to eat %v (meal #%v)\n", p.number, i)

		time.Sleep(time.Second / 2)

		// When a philosopher finishes eating (before it has released its locks)
		// it prints “finishing eating <number>” on a line by itself, where <number>
		// is the number of the philosopher.
		fmt.Printf("finishing eating %v (meal #%v)\n", p.number, i)

		// Note: We let host know that we are done eating.
		// TODO: It would be nice if host can figure it out on his own
		<-h.eating

		p.rightHand.Unlock()
		p.leftHand.Unlock()
	}
}

func main() {

	// There should be 5 philosophers sharing chopsticks,
	// with one chopstick between each adjacent pair of philosophers.
	tableware := make([]*chopstick, 5)
	for i := 0; i < 5; i++ {
		tableware[i] = new(chopstick)
	}

	guests := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		guests[i] = &philosopher{
			number:    i + 1, // Each philosopher is numbered, 1 through 5.
			meals:     3,     // Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
			leftHand:  tableware[i],
			rightHand: tableware[(i+1)%5],
		}
	}

	// The host allows no more than 2 philosophers to eat concurrently.
	h := host{limit: 2}
	h.setLimits()

	t := table{
		host:      &h,
		tableware: tableware,
		chairs:    guests,
	}
	t.feast()
}
