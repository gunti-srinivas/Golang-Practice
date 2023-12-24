package main

import (
	"fmt"
	"math/rand"
)

var x int = 40

func init() {
	fmt.Println("this is from init function")
}
func main() {

	var z int = 2 * rand.Intn(x)
	var x int = rand.Intn(250)
	switch {
	case z < 40:
		fmt.Println("int the first case", z)
	case z > 40:
		fmt.Println("int the second case", z)
	}

	switch {
	case x >= 0 && x <= 50:
		fmt.Println("the value is ", x)
	case x > 50 && x <= 100:
		fmt.Println("the value is ", x)
	case x > 100 && x <= 150:
		fmt.Println("the value is ", x)
	case x > 150 && x <= 200:
		fmt.Println("the value is ", x)
	}
}
