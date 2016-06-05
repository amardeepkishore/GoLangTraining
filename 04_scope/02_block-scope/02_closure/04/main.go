package main

import "fmt"

func amar() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {
	incriment := amar()
	fmt.Println(incriment())
	fmt.Println(incriment())
}
