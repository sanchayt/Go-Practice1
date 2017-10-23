//Package main is given if it is a standalone excutable file
package main

import (
	fm "fmt" //using alias
	rand "math/rand"
)
func main() {

	fm.Println("hello, world") 


	type CNTR int //creating alias for types

	var count CNTR = 10
	fm.Println(count)
	
	//Type Conversion
	a := 45.0
	b := int(a)

	fm.Println(b)


	const c1 = 56/5; //const can be only boolean, number(int, float) or string

	const ( //It can be used as enumeration
		Sunday = iota //iota means give values starting from 0
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	var isOver18 bool = true
	var number int = 678043
	var decimal float64 = 20.077536
	
	var name string = "Sanchay Talwar"

	fm.Printf("Is %s over age of 18? %t and he has $%d and also his age is %.2f\n", name, isOver18, number, decimal )
	
	//Example of loops and switch statements
	i := 0
	for i <= 7 {
		switch i {
		case Sunday: fm.Println("It is Sunday")
		case Monday: fm.Println("It is Monday")
		case Tuesday: fm.Println("It is Tuesday")
		case Wednesday: fm.Println("It is Wednesday")
		case Thursday: fm.Println("It is Thursday")
		case Friday: fm.Println("It is Friday")
		case Saturday: fm.Println("It is Saturday")
	 	default: fm.Println("Wrong Number")
		}

		i++
	}

	//Trying few things with strings
	var test string = "Only for testing purposes!!"

	var x int = rand.Intn(len(test)+1)

	fm.Println("The length is", len(test))
	fm.Println("The ith character is", string(test[x]))
	


}