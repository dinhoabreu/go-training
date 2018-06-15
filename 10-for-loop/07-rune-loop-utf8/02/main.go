package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %q - %v \n", i, string(i), string(i), []byte(string(i)))
	}
}
