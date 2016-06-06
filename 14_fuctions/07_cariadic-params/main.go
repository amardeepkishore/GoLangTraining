package main

import "fmt"

func main() {

	n := average(42, 50, 60, 70, 80)
	fmt.Println("Average of the number passed is ", n)
}

func average(a ...float64) float64 {
	fmt.Println(a)
	fmt.Printf("%T \n", a)

	//total := 0

	var total float64 //it assign value 0 t0 total as above

	for i, v := range a {
		fmt.Println("Index of ", v , "is", i)
		total += v
	}
	return total / float64(len(a))
}