package main

import "fmt"

func main() {
	foo()

	x := func() {
		fmt.Println("I am in anonymous function")
	}

	y := func(s string) {
		fmt.Println(s)
	}

	x()
	y("srinivas")
}

func foo() {
	fmt.Println("I am in foo function")
}

// a named function will have the following
// func (r receiver) identifier (p parameters) return (....){code}

// anonymous function

// func(p parameter(s)) r return(s) {...code...}
