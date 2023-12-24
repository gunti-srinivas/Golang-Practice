package main

import "log"

type print interface {
	Print() string
}

type person struct {
	name string
}

type car struct {
	name string
	base int
}

func (p person) Print() string {
	return p.name
}

func (c car) Print() string {
	return c.name
}

//wrapper function
func Wrapper(logInfo print) {

	log.Println("log from 27 sucessfully processed ", logInfo.Print())
}

func main() {
	var p1 person = person{
		name: "srinivas",
	}

	var c1 car = car{
		name: "volvo",
		base: 20,
	}

	// log.Println(p1.print())

	// log.Println(c1.print())
	Wrapper(p1)

	Wrapper(c1)
}
