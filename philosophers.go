package main

import (
	"fmt"
	"sync"
	"time"
)

type Host struct{ permissions chan int }

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	position        int
	bites           int
}

func (p *Philo) eat(c chan int) {
	fmt.Println("starting to eat ", p.position)
	p.leftCS.Lock()
	p.rightCS.Lock()
	time.Sleep(1 * time.Second)
	p.bites = p.bites - 1
	p.rightCS.Unlock()
	p.leftCS.Unlock()
	fmt.Println("finishing eating ", p.position)
	if p.bites > 0 {
		c <- 1
	} else {
		c <- 0
	}
}

func (h Host) hosting(philos []*Philo) {
	full_ones := 0
	counter := 0
	h.permissions <- 1
	h.permissions <- 1
	for full_ones < len(philos) {
		if philos[counter%5].bites > 0 {
			aux := <-h.permissions
			go philos[counter%5].eat(h.permissions)
			if aux == 0 {
				full_ones = full_ones + 1
			}
		}
		counter = counter + 1
	}
}

func main() {
	var host *Host
	controller := make(chan int, 2)
	host = &Host{controller}
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1, 3}
	}
	host.hosting(philos)
}
