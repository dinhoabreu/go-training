package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// Great usage for assertions
	i := interface{}(b)
	if v, ok := i.(bool); ok {
		fmt.Println(v)
	}

}
