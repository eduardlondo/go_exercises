package main

import (
	"fmt"
	"sync"
)

type chop_stick struct {
}
type Philosopher struct {
	number int
}

func (philosopher Philosopher) eat(ch chan int) {
	defer wait.Done()
	<-ch
	fmt.Println("starting to eat", philosopher.number)
	left := pickChop.Get()
	right := pickChop.Get()
	pickChop.Put(left)
	pickChop.Put(right)
	fmt.Println("finishing eating", philosopher.number)
	ch <- 1
}

var pickChop = sync.Pool{
	New: func() interface{} {
		return new(chop_stick)
	},
}
var wait sync.WaitGroup

func main() {
	ch := make(chan int, 2)
	for i := 0; i < 5; i++ {
		pickChop.Put(new(chop_stick))
	}
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			wait.Add(1)
			go philosophers[i].eat(ch)
		}
	}
	ch <- 1
	ch <- 1
	wait.Wait()
}
