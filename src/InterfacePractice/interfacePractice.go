package main

import (
	"fmt"
	"math"
)
func main() {

	rect1 := Rectangle {10 ,10}
	circ1 := Circle {4}
	fmt.Println("Area of Rectangle is", getArea(rect1))
	fmt.Println("Area of Circle is", getArea(circ1))


}

type Shape interface {
	area() float64
}

type Rectangle struct {
	
	height float64
	width float64
	
}
type Circle struct {
	
	radius float64
	
}
//This a method of struct made
func( r Rectangle) area() float64{
	return r.width * r.height
}
func( c Circle) area() float64{
	return  math.Pi * c.radius * c.radius
}

func getArea(shape Shape) float64 {
	return shape.area()
}