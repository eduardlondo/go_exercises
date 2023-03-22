package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (cow *Cow) Eat() {
	fmt.Println("grass")
}

func (cow *Cow) Move() {
	fmt.Println("walk")
}

func (cow *Cow) Speak() {
	fmt.Println("moo")
}

func (bird *Bird) Eat() {
	fmt.Println("worms")
}

func (bird *Bird) Move() {
	fmt.Println("fly")
}

func (bird *Bird) Speak() {
	fmt.Println("peep")
}

func (snake *Snake) Eat() {
	fmt.Println("mice")
}

func (snake *Snake) Move() {
	fmt.Println("slither")
}

func (snake *Snake) Speak() {
	fmt.Println("hsss")
}

func createAnimal(strs []string) Animal {
	if strs[2] == "cow" {
		return &Cow{strs[1]}
	} else if strs[2] == "bird" {
		return &Bird{strs[1]}
	} else {
		return &Snake{strs[1]}
	}
}

type Animals map[string]*Animal

func main() {
	animals := make(map[string]Animal)
	for true {
		fmt.Print("->")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		strs := strings.Split(input, " ")
		if len(strs) != 3 {
			fmt.Println("invalid input")
			continue
		}
		if strs[0] == "newanimal" {
			animal := createAnimal(strs)
			animals[strs[1]] = animal
			fmt.Println("Created it!")
		} else if strs[0] == "query" {
			animal, ok := animals[strs[1]]
			if !ok {
				fmt.Println("animal not found")
				continue
			}
			switch strs[2] {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("invalid command")
			}
		}
	}

}
