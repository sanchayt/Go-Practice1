package main

import "fmt"

func main() {

	listNums :=	[]float64{42,62,94,16,7}
	
	fmt.Println("Sum :",addUp(listNums))

	num1, num2 := next2Values(5) 
	fmt.Println("num1, num2 :",num1, num2)
	
	fmt.Println(subtractThem(3,4,1,5,7))

	//Closures(function inside main)
	num3 := 6
	doubleValue := func() int {
		num3 *= 2

		return num3
	}
	fmt.Println("double of the number: ",  doubleValue())
	fmt.Println("double of the number: ", doubleValue())


	fmt.Println("factorial of 10: ",factorial(10) )

	fmt.Println(safeDiv(2,0))
	fmt.Println(safeDiv(6,2))
}
//Basic function with one return value
func addUp(numbers []float64) float64{
	sum := 0.0

	for _, val := range numbers{
		sum += val
	}
	return sum
}

//Function with 2 return values
func next2Values(number int) (int, int){
	return number+1, number+2
}

//indefinate number of parameters
func subtractThem(args ...int ) int{
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}
//Recursion example with factorial
func factorial(number int) int {

	if number == 0 {
		return 1
	}

	return number * factorial(number - 1)

}

//defer example
func safeDiv(num1, num2 int) int{
	
	defer func(){
		recover()
	}()

	solution := num1 / num2
	return solution

}