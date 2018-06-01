package main

import "fmt"

const p = "death & taxes"

func main() {
	const q = 42
	fmt.Printf("%v - %T\n", p, p)
	fmt.Printf("%v - %T\n", q, q)
}
