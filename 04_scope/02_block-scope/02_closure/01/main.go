package main

import "fmt"

func main() {
	var x int = 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "This is Gonna rock inside this block only"
		fmt.Println(y)
	}
	//fmt.Println(y)
}
