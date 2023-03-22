package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	m := make(map[string]string)
	fmt.Println("Enter a Name")
	fmt.Scan(&name)
	fmt.Println("Enter an Address")
	fmt.Scan(&address)
	m["name"] = name
	m["address"] = address
	js, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}
