package main

import "fmt"

type person struct {
	name     string
	age      int
	language string
}

type secreatPerson struct {
	entity person
	idx    bool
}

func (p person) speak() {
	fmt.Println(p.language)
}

func (se secreatPerson) speak() {
	fmt.Println("I am secreat person ", se.entity.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		name:     "agent",
		age:      34,
		language: "english",
	}
	p2 := person{
		name:     "vasu",
		age:      45,
		language: "telugu",
	}
	sc := secreatPerson{
		entity: person{
			name:     "secreat",
			age:      45,
			language: "arab",
		},
		idx: true,
	}

	saySomething(p1)
	saySomething(p2)
	saySomething(sc)
}
