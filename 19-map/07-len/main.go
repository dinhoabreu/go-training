package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(len(myGreeting))
	myGreeting["Harleen"] = "Howdy"
	fmt.Println(len(myGreeting))
}
