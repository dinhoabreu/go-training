package main

import (
	"fmt"

	v "github.com/dinhoabreu/go-training/04-scope/01-package/02-visibility/vis" // v as alias for package vis
)

func main() {
	fmt.Println(v.MyName)
	v.PrintNames()
}
