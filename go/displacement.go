package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Please enter acceleration. \n")
	var acceleration float64 
	fmt.Scan(&acceleration)

	fmt.Printf("Please enter initial velocity. \n")
	var initialVelocity float64
	fmt.Scan(&initialVelocity)

	fmt.Printf("Please enter initial displacement. \n")
	var initialDisplacement float64
	fmt.Scan(&initialDisplacement)

	fmt.Printf("Please enter time. \n")
	var time float64
	fmt.Scan(&time)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println(fn(time))
}

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return 0.5 * acceleration * math.Pow(time, 2) + initialVelocity * time + initialDisplacement
	}
}