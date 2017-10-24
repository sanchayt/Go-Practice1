package main

import "fmt"

func main() {

	x := 0

	changeXVal(x)
	fmt.Println("x =", x)

	changeXValNow(&x)
	
	fmt.Println("x =", x)
	fmt.Println("Memory Address for x =", &x)

	yPtr := new(int)
	changeYValNow(yPtr)

	
	fmt.Println("y =", *yPtr)
	
}

func changeXVal( x int){
	
	x = 2

}
func changeXValNow( x *int){
	
	*x = 2

}
func changeYValNow( yPtr *int){
	
	*yPtr = 200

}