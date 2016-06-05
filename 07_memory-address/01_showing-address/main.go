package main

import "fmt"

func main() {

	a := 42

	fmt.Println("a - ", a)
	fmt.Println("a's memory address is ", &a)
	//showing decimal value of a
	fmt.Printf("%d \n", &a)
	//showing binary value of a
	fmt.Printf("%b \n", &a)

}
