package main

import "fmt"

func main() {
	recursion(10)
}

func recursion(i int) {
	if i == -1 {
		return
	}
	fmt.Println(i)
	recursion(i - 1)
	return
}
