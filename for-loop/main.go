package main

import "fmt"

func main() {
	var y int = 5

	for i := 1; i < 10; i++ {
		fmt.Printf("the value is %v \n", i)
	}

	// like while loop in other languages
	for y < 20 {
		fmt.Printf("the value is %v \n", y)
		y++
	}

	//using continue in the for loop

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("third for loop ", i)
	}
}
