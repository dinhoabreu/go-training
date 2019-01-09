package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(t1.Format("03:04"))
	fmt.Println(t1.Format("01 02 2006 03:04"))
	fmt.Println(t1.Format(time.RFC3339))

	t2 := time.Now().Add(2 * time.Minute)
	fmt.Println(t2.Format("03:04"))

	fmt.Println(t1.Before(t2))
}
