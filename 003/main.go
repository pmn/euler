/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import (
	"fmt"
	"math"
)

func sieve_primes(n int) []int {
	var candidates []int
	candidates = append(candidates, 2)

	// Build a naive list of integers
	allnums := make([]bool, n)
	for i := range allnums {
		allnums[i] = true
	}
	// Set initial primality for 0, 1, 2
	allnums[0] = false
	allnums[1] = false

	// Quick and dirty implementation of sieve of eratosthenes
	for i := range allnums {
		if allnums[i] {
			candidates = append(candidates, i)
			for j := i; j*i < n; j++ {
				if i*j < n-1 {
					allnums[i*j] = false
				}
			}
		}
	}

	var primes []int
	for i := range allnums {
		if allnums[i] && i < n {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	test := 600851475143
	lbound := math.Ceil(math.Sqrt(float64(test)))

	candidates := sieve_primes(int(lbound))
	var primes []int

	for i := range candidates {
		if test%candidates[i] == 0 {
			primes = append(primes, candidates[i])
		}
	}

	fmt.Println(primes)
}
