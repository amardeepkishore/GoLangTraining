package main

import "fmt"

func zero(a *int) {
	*a = 0

}

func main() {
	b := 5
	zero(&b)
	fmt.Println(b)

}
