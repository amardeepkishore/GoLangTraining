package main

import "fmt"

const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

// Reset the iota value an start from 0 again
const (
	D = iota // 0
	E = iota // 1
	F = iota // 2
)

// Ignore the 0 and continue

const (
	_ = iota // 0
	G = iota // 1
	H = iota // 2
)

// bit shifting or bitwise operstion << or >>

const (
	_ = iota //0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Printf("%b \t", KB)
	fmt.Printf("%d \n", KB)
	fmt.Printf("%b \t", MB)
	fmt.Printf("%d \n", MB)
	fmt.Printf("%b \t", GB)
	fmt.Printf("%d \n", GB)
}
