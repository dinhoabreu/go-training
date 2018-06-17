package main

import (
	"fmt"
	"strings"
)

type slicebytes []byte

func (sb slicebytes) String() string {
	out := make([]string, 0, 4)
	for _, v := range sb {
		out = append(out, fmt.Sprintf("%08b", v))
	}
	return "[" + strings.Join(out, " ") + "]"
}

func main() {
	var x [256]string

	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 256; i++ {
		x[i] = string(i)
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %v - %v\n", v, v, []byte(v), slicebytes(v))
	}
}
