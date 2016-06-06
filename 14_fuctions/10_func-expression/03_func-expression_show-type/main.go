package main

import "fmt"

func main() {

	greating := func() {
		fmt.Println("Hello World !")
	}
	greating()
	fmt.Printf("%T \n", greating)
}
