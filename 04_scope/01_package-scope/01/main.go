package main

import "fmt"

//package level scope, can be called anywhere in this package (01). NOTE:- package is not main but 01, main is just making 01 as main package.


var x int = 42

func main() {
	fmt.Printf("%v \n", x)
	foo()
}

func foo() {
	fmt.Printf("%v \n", x)
}
