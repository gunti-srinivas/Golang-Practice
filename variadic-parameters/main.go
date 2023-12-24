package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}

	//first way of calling the sum is
	// var totalsum = sum("sai", 1, 2, 3, 4, 5, 6)

	//second way of calling the sum
	var totalsum int = sum("sai", xi...)
	fmt.Println(totalsum)
}

// what happens is we get slice of integers and we go will create the slice datastructure
// using range for loop
//variadic parameter should be final incoming parameter
func sum(s string, ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T \n", ii)
	var sum int = 0
	for i, val := range ii {
		fmt.Printf("The index is %d and the value is %d \n ", i, val)
		sum = sum + val
	}
	return sum
}

//func (r receiver) identifier (parameters) (return (s)){ code }
