package main

import "fmt"

var arg float64
var argint int

func main() {
	fmt.Println("Write a float")
	fmt.Scan(&arg)
	argint = int(arg)
	fmt.Println(argint)
}
