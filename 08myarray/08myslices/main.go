package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to a slices")

	var fruitList = []string{"Apple", "Orange"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Papaya")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 2234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 4

	fmt.Println(highScores)
	highScores = append(highScores, 55, 65)
	fmt.Println(highScores)

	//remove values from slices

	var courses = []string{"reactjs", "javascript", "ruby", "java"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1])
	fmt.Println(courses)

}
