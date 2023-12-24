package main

import "fmt"

type person struct {
	first         string
	age           int
	qualification string
}

func (p person) speak() {
	fmt.Println(p.first, p.age, p.qualification)
}

func main() {

	p1 := person{
		first:         "james",
		age:           39,
		qualification: "btech",
	}

	p2 := person{
		first:         "kalyan",
		age:           23,
		qualification: "masters",
	}

	p3 := person{
		first: "sai",
		age:   33,
	}
	fmt.Println("---checking---")
	p1.speak()
	p2.speak()
	p3.speak()
	fmt.Println("-----> ", p3.qualification)

}
