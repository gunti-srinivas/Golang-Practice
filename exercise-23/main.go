package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 1; i <= 100; i++ {
		var x int = rand.Intn(10)
		var y int = rand.Intn(10)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("%v and %v are less than 4 \n", x, y)
		case x > 6 && y > 6:
			fmt.Printf("%v and %v are less than 6 \n", x, y)
		default:
			fmt.Println("none of the previous conditions satisfied")
		}
	}
}
