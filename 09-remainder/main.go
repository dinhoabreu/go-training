package main

import "fmt"

func main() {
	x := 13 % 3
	if x == 1 {
		fmt.Println("1 Odd")
	} else {
		fmt.Println("1 Even")
	}

	for i := 1; i < 7; i++ {
		if i%2 == 1 {
			fmt.Printf("%d Odd\n", i)
		} else {
			fmt.Printf("%d Even\n", i)
		}
	}
}
