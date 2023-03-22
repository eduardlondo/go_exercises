package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	counter++
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	time.Sleep(time.Second)
	fmt.Println(counter)
}

/**
The increment() function simply increments a global variable called counter.
The main function starts 1000 goroutines that each call the increment() function.
After all the goroutines have had a chance to run, the main function prints out the value of counter.
The race condition in this code occurs because the counter variable is accessed concurrently by multiple goroutines.
if two goroutines try to increment counter at the same time,
they may both read the same value of counter and then write back the same value,
effectively only incrementing the counter once.
**/
