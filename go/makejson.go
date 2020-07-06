package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Please enter your name.\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Please enter your address.\n")
	scanner.Scan()
	address := scanner.Text()

	myMap := make(map[string]string)
	myMap["name"] = name
	myMap["address"] = address

	jsonObject, _ := json.Marshal(myMap)
	fmt.Printf(string(jsonObject))
}
