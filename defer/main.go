// func (receiver) identifier (parameters) return (types) { code }

package main

import "fmt"

func main() {
	defer foo()
	boo()
	boo1()
	boo2()
	boo3()
}

func foo() {
	fmt.Println("I am executing last")
}

func boo() {
	fmt.Println("I am done")
}

func boo1() {
	fmt.Println("I am in boo1")
}

func boo2() {
	fmt.Println("I am in boo2")
}
func boo3() {
	fmt.Println("I am in boo3")
}
