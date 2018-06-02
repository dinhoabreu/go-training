package main

import "fmt"

const (
	_ = iota // 0
	// KB Kilo bytes
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	// MB Mega bytes
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	// GB Giga bytes
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	// TB Tera bytes
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
