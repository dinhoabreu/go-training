package main

import "fmt"

func main() {
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))
	fmt.Println(string([]byte{'O', 'l', 'á'}))
	fmt.Println(string([]rune{'O', 'l', 'á'}))
	// conversion: []bytes to string
	// we'll learn about []bytes soon
}
