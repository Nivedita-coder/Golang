package main

import "fmt"

func main() {
	fmt.Println("Welcoem to a class on pointer")

	//var ptr *int
	//fmt.Println("Value of ponter is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}
