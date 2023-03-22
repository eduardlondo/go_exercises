package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	sli := make([]Name, 0)
	fmt.Scan(&fileName)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// do something with a line
		names := strings.Split(scanner.Text(), " ")
		aux := Name{fname: names[0], lname: names[1]}
		sli = append(sli, aux)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, v := range sli {
		fmt.Println("First name: " + v.fname + " And second name: " + v.lname)
	}
}
