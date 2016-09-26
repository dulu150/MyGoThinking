package main

import "fmt"

type PersonInfo struct {
	ID   string
	Name string
}

func main() {
	var Person map[string]PersonInfo

	Person = make(map[string]PersonInfo)

	Person["1"] = PersonInfo{"nickel zhang", "hello the world"}

	person, ok := Person["1"]
	if ok {
		fmt.Println("result:", person.ID)
	} else {
		fmt.Println("failed")
	}
}
