package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a  //* is part of type of value, where int is stored in memory address.

	fmt.Println(b)
}
