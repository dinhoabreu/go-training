package main

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("hello")
	}
}

// run this at command:
// go test -bench='.*'
