package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 42; i++ {
		var x int = rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("The value is %d \n", x)
		case 1:
			fmt.Printf("The value is %d \n", x)
		case 2:
			fmt.Printf("The value is %d \n", x)
		case 3:
			fmt.Printf("The value is %d \n", x)
		case 4:
			fmt.Printf("The value is %d \n", x)
		}
	}

}
