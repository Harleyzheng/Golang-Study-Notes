package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := map[string]Animal{}

	for {
		fmt.Printf(">\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()
		requestArray := strings.Fields(request)
		name := requestArray[1]
		switch requestArray[0] {
		case "newanimal":
			animalType := requestArray[2]
			var animal Animal

			switch animalType {
			case "cow":
				animal = Cow{"grass", "walk", "moo"}
				animals[name] = animal
			case "bird":
				animals[name] = Bird{"worms", "fly", "peep"}
			case "snake":
				animals[name] = Snake{"mice", "slither", "hsss"}
			default:
				fmt.Println("Error animal type")
			}
			fmt.Printf("Created it!")

		case "query":
			animal, present := animals[name]
			move := requestArray[2]
			if present {
				switch move {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				default:
					fmt.Println("Error move type")
				}
			} else {
				fmt.Printf("Animal " + name + " does not exist!")
			}

		default:
			fmt.Printf("Error. Please try again \n")
		}
	}
}
