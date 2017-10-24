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
	
	for i := 0; i <= 7; i++ {
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

	}

	//Trying few things with strings
	var test string = "Only for testing purposes!!"

	var x int = rand.Intn(len(test)+1)

	fm.Println("The length is", len(test))
	fm.Println("The ith character is", string(test[x]))
	
	//Arrays
	var favNums[5] float64

	favNums[0] = 12.3
	favNums[1] = 135
	favNums[2] = 405.7
	favNums[3] = 69.29
	favNums[4] = 2.0486

	fm.Println(favNums[3])

	favNums2 := [5]float64 {68,98,12,39,20}
	//with index
	for i, value := range favNums2 {
		fm.Println(value, i)
	}
	//without index
	for _, value := range favNums2 {
		fm.Println(value)
	}
	
	//Slice example
	numSlice := []int {5,4,3,2,1}

	numSlice2 := numSlice[3:5]
	fm.Println("numSlice2[2] =",numSlice2[2])

	fm.Println("numSlice[:2] =",numSlice[:2])

	
	
}