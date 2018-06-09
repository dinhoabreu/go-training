package main

import "fmt"

func increasePrimeFactors(primeFactors []int) (int, []int) {
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
			return num, primeFactors
		}
	}
}

func getPrimeFactors(n int, primeFactors []int) ([]int, []int) {
	r := n
	factors := []int{}
	checkAndReduce := func(primeFactor int) bool {
		if n%primeFactor == 0 {
			r /= primeFactor
			factors = append(factors, primeFactor)
			return r == 1
		}
		return false
	}
	for _, primeFactor := range primeFactors {
		if checkAndReduce(primeFactor) {
			return factors, primeFactors
		}
	}
	primeFactor := primeFactors[len(primeFactors)-1]
	for primeFactor < n {
		primeFactor, primeFactors = increasePrimeFactors(primeFactors)
		if checkAndReduce(primeFactor) {
			return factors, primeFactors
		}
	}
	return factors, primeFactors
}

func printResult(n int, factors []int) {
	fmt.Printf("The prime factors of %d are %v\n", n, factors)
	fmt.Printf("So, the largest prime factor of the number %d is %d\n", n, factors[len(factors)-1])
	fmt.Println()
}

func main() {
	primeFactors := []int{2, 3, 5}
	for _, n := range []int{13195, 600851475143} {
		var factors []int
		factors, primeFactors = getPrimeFactors(n, primeFactors)
		printResult(n, factors)
	}
}

/*

[Largest prime factor](https://projecteuler.net/problem=3)
Problem 3
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/
