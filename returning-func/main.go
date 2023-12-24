package main

import "fmt"

func main() {

	y := foo()

	fmt.Println(y)

	x := boo()

	fmt.Println(x())
}

func foo() int {
	return 43
}

func boo() func() int {
	return func() int {
		return 56
	}
}
