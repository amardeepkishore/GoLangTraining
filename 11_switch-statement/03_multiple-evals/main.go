package main

import "fmt"

func main() {

	switch "Amar" {

	case "Seema":
		fmt.Println("Seema is wife.")
	case "Amar", "Mac"://Multiple evals
		fmt.Println("Found Amar and Mac.")
	case "Neeraj":
		fmt.Println("At Home")
	default:
		fmt.Println("Nothing Found")
	}
}