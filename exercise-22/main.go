package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i int = rand.Intn(10)
	var j int = rand.Intn(10)
	fmt.Println(i)
	fmt.Println(j)
	if i < 4 && j < 4 {
		fmt.Println("i and j are less than 4")
	} else if i > 6 && j > 6 {
		fmt.Println("i and j greater than 6")
	} else {
		fmt.Println("none of the previous conditions are satisfied")
	}
}
