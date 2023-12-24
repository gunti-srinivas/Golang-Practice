package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	// short hand deceleration
	x := 70
	y := 50

	fmt.Printf(" x = %v \n  y = %v \n", x, y)

	if x < 40 {
		fmt.Println("x is less than 40")
	} else if x == 40 {
		fmt.Println("x is equal to 40")
	} else {
		fmt.Println("x is greater than 40")
	}
	if x > 30 && x < 50 {
		fmt.Println("x is inbetween 30 & 40")
	} else if x > 60 && x < 90 {
		fmt.Println("x is inbetween 60 and 90")
	}

	//statememt statement idiom

	if z := 2 * rand.Intn(x); z <= x {
		fmt.Printf("The value of z is %v for the x %v ", z, x)
	} else {
		fmt.Printf("The value of z is %v for the x %v ", z, x)
	}
}
