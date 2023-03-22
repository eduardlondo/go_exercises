package main

import (
	"fmt"
	"strings"
)

var arg string
var x, y, z, found bool

func main() {
	fmt.Println("Write something")
	fmt.Scan(&arg)
	arg = strings.ToLower(arg)
	x = strings.HasPrefix(arg, "i")
	y = strings.HasSuffix(arg, "n")
	z = strings.Contains(arg, "a")
	found = x && y && z
	if found {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
