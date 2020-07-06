package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.ToLower(s)

	condition := strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Index(s, "a") != -1

	if condition {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
