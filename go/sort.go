package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strings"
	"strconv"
	"sync"
)

func main() {
	fmt.Printf("Please enter a series of numbers as a string, separated by space: \n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	intStringArray := strings.Fields(line)
	slice := make([]float64, 0, 3)

	for _, value := range intStringArray {
		intValue, _ := strconv.Atoi(value)
		slice = append(slice, float64(intValue))
	}

	slice1 := make([]float64, 0, 3)
	slice2 := make([]float64, 0, 3)
	slice3 := make([]float64, 0, 3)
	slice4 := make([]float64, 0, 3)

	index := 0
	for index < len(slice)/4 {
		slice1 = append(slice1, slice[index])
		index++
	}

	for index >= len(slice)/4 && index < 2*len(slice)/4 {
		slice2 = append(slice2, slice[index])
		index++
	}

	for index >= 2*len(slice)/4 && index < 3*len(slice)/4 {
		slice3 = append(slice3, slice[index])
		index++
	}

	for index >= 3*len(slice)/4 && index < len(slice) {
		slice4 = append(slice4, slice[index])
		index++
	}

	var wg sync.WaitGroup
	wg.Add(4)
	go sort(slice1, &wg)
	go sort(slice2, &wg)
	go sort(slice3, &wg)
	go sort(slice4, &wg)
	wg.Wait()

	merged := merge(slice1, slice2, slice3, slice4)

	fmt.Printf("%v", merged)
}

func merge(s1 []float64, s2 []float64, s3 []float64, s4 []float64) []float64 {
	merged := make([]float64, 0, 3)
	i1 := 0
	i2 := 0
	i3 := 0
	i4 := 0

	for i1 < len(s1) || i2 < len(s2) || i3 < len(s3) || i4 < len(s4) {
		next := float64(MaxUint >> 1) 
		if i1 < len(s1) {
			next = math.Min(next, s1[i1])
		}
		
		if i2 < len(s2) {
			next = math.Min(next, s2[i2])
		}

		if i3 < len(s3) {
			next = math.Min(next, s3[i3])
		}

		if i4 < len(s4) {
			next = math.Min(next, s4[i4])
		}

		if i1 < len(s1) && next == s1[i1] {
			i1++
		} else if i2 < len(s2) && next == s2[i2] {
			i2++
		} else if i3 < len(s3) && next == s3[i3] {
			i3++
		} else if i4 < len(s4) && next == s4[i4] {
			i4++
		}

		merged = append(merged, next)
	}

	return merged
}

func swap(items []float64, i int, j int) {
	items[i+1], items[i] = items[i], items[i+1]
}

const MaxUint = ^uint(0) 

func sort(items []float64, wg *sync.WaitGroup) {
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

	wg.Done()
}
	