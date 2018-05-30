package main

import "fmt"

func main() {
	// fmt.Println(x) // x at this point is undefined
	fmt.Println(y)
	x := 42
	fmt.Println(x)
}

var y = 42
