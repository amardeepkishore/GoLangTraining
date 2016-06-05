package main

import "fmt"

var x = 0

func incriment() int {
	x++
	return x
}

func main() {
	fmt.Println(incriment())
	fmt.Println(incriment())
}
