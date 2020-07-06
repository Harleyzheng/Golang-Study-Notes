package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var slice = make([]int, 0, 3)

	for {
		fmt.Printf("Please enter a number or X to terminate.")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		if "X" == s {
			fmt.Printf("Terminated")
			break
		}

		i, err := strconv.Atoi(s)

		if err != nil {
			fmt.Printf("Invalid input")
			break
		}

		slice = append(slice, i)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}
