package main

import "fmt"

func main() {
	g := incremented()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

	f := incremented()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func incremented() func() (int, int) {
	var x int = 0
	return func() (int, int) {
		y:= 0
		y++
		x++
		return (x, y)
	}
}
