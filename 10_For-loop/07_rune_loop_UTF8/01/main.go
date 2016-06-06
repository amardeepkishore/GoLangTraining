package main

import "fmt"

func main() {

	for i := 100; i <= 200; i++ {

		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))

		//fmt.Println('i') // it print 105, because in unicode chart value of i is 105
	}
}


//rune is a character, alias for int32, is a 4 byte UTF-8 (string is a collection of rune)
//each rute is 1 to 4 byte