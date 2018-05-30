package main

import "fmt"

func main() {
	var a int
	var b string
	var c float32
	var d bool
	var e []rune

	fmt.Printf("%v\t%#v\t%T\n", a, a, a)
	fmt.Printf("%v\t%#v\t%T\n", b, b, b)
	fmt.Printf("%v\t%#v\t%T\n", c, c, c)
	fmt.Printf("%v\t%#v\t%T\n", d, d, d)
	fmt.Printf("%v\t%#v\t%T\n", e, e, e)
}
