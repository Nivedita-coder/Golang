package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	nivedita := User{"Nivedita", "nivedita@go.dev", true, 20}
	fmt.Println(nivedita)
	fmt.Printf("nivedita details are: %+v\n", nivedita)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
