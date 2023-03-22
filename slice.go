package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0)
	for true {
		var input string
		/* var numberValue int */
		fmt.Println("Enter a number, enter X for close the program ")
		fmt.Scan(&input)
		if input == "x" || input == "X" {
			break
		} else {
			if numberValue, err := strconv.Atoi(input); err == nil {
				sli = append(sli, numberValue)
				sort.Ints(sli)
				fmt.Print("[ ")
				for i, v := range sli {
					fmt.Print(v)
					if i != (len(sli) - 1) {
						fmt.Print(" , ")
					}
				}
				fmt.Println("]")
			}
		}
	}
}
