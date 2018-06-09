package main

import "fmt"

var primeFactors = []int{2, 3, 5}

func increasePrimeFactors() int {
	num := primeFactors[len(primeFactors)-1]
	for {
		num++
		isPrimeFactor := true
		for _, primeFactor := range primeFactors {
			if num%primeFactor == 0 {
				isPrimeFactor = false
				break
			}
		}
		if isPrimeFactor {
			primeFactors = append(primeFactors, num)
			return num
		}
	}
}

func primeFactorsIterator() func() int {
	index := 0
	return func() int {
		if index >= len(primeFactors) {
			increasePrimeFactors()
		}
		factor := primeFactors[index]
		index++
		return factor
	}
}

func getPrimeFactors(n int) []int {
	r := n
	factors := []int{}
	next := primeFactorsIterator()
	for {
		primeFactor := next()
		if n%primeFactor == 0 {
			r /= primeFactor
			factors = append(factors, primeFactor)
			if r == 1 {
				return factors
			}
		}
	}
}

func printResult(n int, factors []int) {
	fmt.Printf("The prime factors of %d are %v\n", n, factors)
	fmt.Printf("So, the largest prime factor of the number %d is %d\n", n, factors[len(factors)-1])
	fmt.Println()
}

func main() {
	for _, n := range []int{13195, 600851475143} {
		var factors []int
		factors = getPrimeFactors(n)
		printResult(n, factors)
	}
}

/*

[Largest prime factor](https://projecteuler.net/problem=3)
Problem 3
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/
