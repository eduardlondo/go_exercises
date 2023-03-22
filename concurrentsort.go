package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func bubbleSort(wg *sync.WaitGroup, nums []int) {
	/* fmt.Println(nums) */
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j)
			}
		}
	}
	fmt.Println("subarray sorted ", nums)
	wg.Done()
}

func swap(nums []int, index int) {
	temp := nums[index]
	nums[index] = nums[index+1]
	nums[index+1] = temp
}

func readIntegers() []string {

	fmt.Println("Enter more tan 3 integers, separated by spaces:")

	// Read the input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strs := strings.Split(input, " ")
	return strs
}

func main() {
	fmt.Printf("Core availables are %d\n", getGOMAXPROCS())
	strs := readIntegers()
	var wg sync.WaitGroup
	var nums []int
	var base_size int
	counter := 0
	slices := make([][]int, 4)
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	if len(nums) < 4 {
		fmt.Println("hay menos de 4 numeros")
	} else {
		base_size = len(nums) / 4
		module := len(nums) % 4
		for i := range slices {
			cols := base_size
			if counter < module {
				cols = base_size + 1
				counter++
			}
			slices[i] = make([]int, cols)
		}
		counter = 0
		for i := range slices {
			for j := range slices[i] {
				slices[i][j] = nums[counter]
				counter++
			}
		}
		wg.Add(4)
		for i := range slices {
			go bubbleSort(&wg, slices[i])
		}
		wg.Wait()

		ab := merge(slices[0], slices[1])
		cd := merge(slices[2], slices[3])
		result := merge(ab, cd)
		fmt.Print("all sorted as ")
		fmt.Println(result)
	}

}
