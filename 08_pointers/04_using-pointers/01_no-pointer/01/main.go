package main

import "fmt"

func zero(a int) {
	a = 0
}

func main() {
	b := 5
	zero(b) // fucntio zero has its own memory address so fmt.Println(5) remains 5.
	fmt.Println(b)
}
