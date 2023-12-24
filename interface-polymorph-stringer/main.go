package main

import "fmt"

type person struct {
	name string
}

func (p person) getName() string {
	return p.name
}

func (p person) string() string {
	return p.getName()
}

func main() {
	var p1 person = person{
		name: "srinivas",
	}

	var p2 person = person{
		name: "krishna",
	}

	var p3 person = person{
		name: "basha",
	}

	fmt.Println(p1.string())

	fmt.Println(p2.string())

	fmt.Println(p3.string())
}
