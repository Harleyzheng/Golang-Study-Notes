package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	number int
}

func (p *Philosopher) Eat(chops []chan int, host chan int, wg *sync.WaitGroup) {
	first, second := getChops(p.number)
	c := <-host
	select {
	case <-chops[first]:
		<-chops[second]
	case <-chops[second]:
		<-chops[first]
	}

	fmt.Printf("starting to eat %d\n", p.number)

	time.Sleep(1 * time.Second)

	fmt.Printf("finishing eating %d\n", p.number)

	select {
	case chops[first] <- 1:
		chops[second] <- 1
	case chops[second] <- 1:
		chops[first] <- 1
	}
	host <- c
	wg.Done()
}

func getChops(number int) (int, int) {
	switch number {
	case 0:
		return 1, 4
	case 1:
		return 0, 2
	case 2:
		return 1, 3
	case 3:
		return 2, 4
	case 4:
		return 3, 0
	default:
		fmt.Println("Error")
		return 0, 0
	}
	return 0, 0
}

func main() {
	// initialize chopsticks
	var chops []chan int
	for i := 0; i < 5; i++ {
		chop := make(chan int, 1)
		chop <- 1
		chops = append(chops, chop)
	}

	// initialize host
	host := make(chan int, 2)
	host <- 0
	host <- 1

	// initialize philosophers
	var philosophers []Philosopher
	for i := 0; i < 5; i++ {
		philosophers = append(philosophers, Philosopher{i})
	}

	var wg sync.WaitGroup
	wg.Add(15) // increments the counter

	// start eating
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			go philosophers[j].Eat(chops, host, &wg)
		}
	}

	wg.Wait() // blocks until counter = 0

}
