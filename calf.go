package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// NINTS number of integers
	NINTS = 10
)

// Swap swap slice element i and i+1
func Swap(pair []int, i int) {
	pair[i], pair[i+1] = pair[i+1], pair[i]
}

// BubbleSort naive implementation of bubble sort
func BubbleSort(ints []int) {
	n := len(ints)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if ints[j] > ints[j+1] {
				Swap(ints, j)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ints := []int{}

	for i := 0; i < NINTS; i++ {
		text, _ := reader.ReadString('\n')
		v, _ := strconv.Atoi(strings.Trim(text, "\n"))
		ints = append(ints, v)
	}
	// fmt.Println(ints)
	BubbleSort(ints)
	fmt.Println(ints)
}
