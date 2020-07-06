package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var slice = make([]int, 0, 5)

	fmt.Printf("Please provide a sequence of integars.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	intString := scanner.Text()
	intStringArray := strings.Fields(intString)

	for _, value := range intStringArray {
		intValue, _ := strconv.Atoi(value)
		slice = append(slice, intValue)
	}

	bubblesort(slice)
	fmt.Println(slice)
}

func swap(items []int, i int, j int) {
	items[i+1], items[i] = items[i], items[i+1]
}

func bubblesort(items []int) {
	n := len(items)
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				swap(items, i, i+1)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
	