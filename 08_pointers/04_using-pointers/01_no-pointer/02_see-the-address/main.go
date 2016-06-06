package main

import "fmt"

func zero(a int) {
	fmt.Println(&a)
	a = 0
}

func main() {
	b := 5
	fmt.Println(&b)
	zero(b) // fucntio zero has its own memory address so fmt.Println(5) remains 5.
	fmt.Println(b)
}
