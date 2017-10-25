package main

import "fmt"
func main() {

	rect1 := Rectangle {0, 50, 10 ,10}
	fmt.Println("Rectangle is", rect1.width, "wide")
	fmt.Println("Area of Rectangle is", rect1.area())
	
}


//This is a struct
type Rectangle struct {

	leftX float64
	topY float64
	height float64
	width float64

}
//This a method of struct made
func( rect *Rectangle) area() float64{
	return rect.width * rect.height
}
