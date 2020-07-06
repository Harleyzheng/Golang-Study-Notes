package main

import (
	"fmt"
	"time"
)

var result int 

func plusOne () {
	result = result + 1
}

func double () {
	result = result * 2
}

/*
	The outcome of running this code is indeterministic because the goroutines are executed concurrently.
	There's a race condition.
*/
func main() {
	result = 1
	go plusOne()
	go double()
	time.Sleep(3000)
	fmt.Println(result)
}
