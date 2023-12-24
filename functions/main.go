package main

import "fmt"

func main() {
	foo()
	bar("jhon")
	var val string = aloha("T")
	fmt.Println(val)
	value1, value2 := dogYears("srinivas", 67)
	fmt.Println(value1, value2)
}

func foo() {
	fmt.Println("I am in foo")
}

func bar(s string) {
	fmt.Println("My name is ", s)
}

func aloha(s string) string {
	return fmt.Sprintf("They call me Mr. %v", s)
}

// two parameters and two return types
func dogYears(s string, i int) (string, int) {
	return s, i
}
