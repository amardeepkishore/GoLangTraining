package main

import "fmt"

func main() {

	switch "Seema" {

	case "Seema":
		fmt.Println("Seema is wife.")
		fallthrough
	case "Amar":
		fmt.Println("Found Amar.")
	default:
		fmt.Println("Nothing Found")
	}
}