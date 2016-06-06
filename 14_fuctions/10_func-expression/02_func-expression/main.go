package main

import "fmt"

func main() {

	greating := func() {
		fmt.Println("Hello World !") // assinging a anonymous (No Name)function is called func expression.
	}				// The only way to define a function inside a function.

	greating()
}
