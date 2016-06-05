package main

import "fmt"

//function without a name, creating a function and assigning it to a variable.

// func() is anonymous function
// incriment := func() is function expression.

//Only way to define fuction inside a fuction.

func main() {

	x := 0
	incriment := func() int {
		x++
		return x
	}

	fmt.Println(incriment())
	fmt.Println(incriment())

}
