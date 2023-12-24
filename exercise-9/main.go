package main

import ("fmt")

var x int = 40
var y int = 41

func main(){
	z := 42
	fmt.Printf("The value of x is %v and the type is %T \n", x, x);
	fmt.Printf("The value of z is %v and the type is %T \n", z, z);
	fmt.Printf("The value of y is %v and the type is %T \n", y, y);
}