package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	structs := make([]Person, 0, 10)

	fmt.Printf("Please enter file name.\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()

	file, _ := os.Open(fileName)
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		fullName := scanner.Text()
		names := strings.Fields(fullName)
		structs = append(structs, Person{names[0], names[1]})
	}

	for _, value := range structs {
		fmt.Printf(value.fname + " " + value.lname + "\n")
	}
}
