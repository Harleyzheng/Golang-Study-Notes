package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func main() {
	animals := map[string]Animal {
		"cow": Animal{"grass", "walk", "moo"},
		"bird": Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		fmt.Printf("Please enter request. >\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()
		requestArray := strings.Fields(request)
		animal := animals[requestArray[0]]
		switch requestArray[1] {
		case "eat":
			fmt.Printf(animal.Eat() + "\n")

		case "move":
			fmt.Printf(animal.Move() + "\n")

		case "speak":
			fmt.Printf(animal.Speak() + "\n")

		default:
			fmt.Printf("Error. Please try again \n")
		}
	}
}
