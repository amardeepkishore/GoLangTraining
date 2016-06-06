package main

import "fmt"

func zero(a *int) {
	fmt.Println("Memory Address of a", a)
	*a = 0

}

func main() {
	b := 5
	fmt.Println("Memory Address of b", &b)
	zero(&b)
	fmt.Println(b)

}