package main

import "fmt"

func main() {

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Printf("the outer loop is at %v    inner loop is at %v \n", i, j)
		}
	}
}
