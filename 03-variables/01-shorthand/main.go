package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := []rune(b)

	fmt.Printf("%v\t%#v\t%T\n", a, a, a)
	fmt.Printf("%v\t%#v\t%T\n", b, b, b)
	fmt.Printf("%v\t%#v\t%T\n", c, c, c)
	fmt.Printf("%v\t%#v\t%T\n", d, d, d)
	fmt.Printf("%v\t%#v\t%T\n", e, e, e)
}
