package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var n int
	var pids []string
	if !parseArgs(&n, &pids) {
		os.Exit(1)
	}
	result := make(chan int)
	for _, pid := range pids {
		go func(pid string) {
			result <- incrementor(pid, n)
		}(pid)
	}
	count := 0
	for range pids {
		count += <-result
	}
	fmt.Println("Final Counter:", count)
}

func incrementor(pid string, n int) int {
	for i := 0; i < n; i++ {
		fmt.Printf("Process: %v printing: %v\n", pid, i)
	}
	return n
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] pid1 [pid2] [...]\n", os.Args[0])
	flag.PrintDefaults()
}

func parseArgs(n *int, pids *[]string) bool {
	nFlag := flag.Int("n", 20, "count until 'n'")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return false
	}
	*n = *nFlag
	*pids = flag.Args()
	return true
}
