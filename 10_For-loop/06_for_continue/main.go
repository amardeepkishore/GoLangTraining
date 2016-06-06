package main

import "fmt"

func main() {
	i := 0

	for {
		i++
		if i%2 == 0 {
			fmt.Println("Even -", i)
			continue
		}

		fmt.Println("Odd - ", i)

		if i >= 50 {
			break
		}
	}

}
