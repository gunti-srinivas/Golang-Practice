package main

import "fmt"

func main() {
	//data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Println("ranging over slice ", i, v)
	}

	//for range loop

	// data structure - map

	m := map[string]int{
		"key1": 45,
		"key2": 56,
	}

	for k, v := range m {
		fmt.Println("ranging over map ", k, v)
	}
}
