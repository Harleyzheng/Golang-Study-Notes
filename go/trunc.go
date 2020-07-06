package main

import "fmt"

func main() {
	var x float32
	for {
		fmt.Scan(&x)
		fmt.Printf("%d", int(x))
	}
}
