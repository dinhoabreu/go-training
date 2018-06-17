package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// food is out of scope
	// fmt.Println(food)

}
