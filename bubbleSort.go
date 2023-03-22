package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	strs := readIntegers()

	var nums []int
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	bubbleSort(nums)
	// Print the sorted integers
	fmt.Println("Sorted integers:")
	for i, num := range nums {
		fmt.Print(num)
		if i < 10-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// Convert the input strings to integers
func readIntegers() []string {

	fmt.Println("Enter up to 10 integers, separated by spaces:")

	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strs := strings.Split(input, " ")
	return strs
}

// Use Bubble Sort to sort the integers
func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j)
			}
		}
	}
}

// Swap the values
func swap(nums []int, index int) {
	temp := nums[index]
	nums[index] = nums[index+1]
	nums[index+1] = temp
}
