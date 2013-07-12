// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?

package main

import "fmt"

func find_nth_prime(n int) int {
	var primes []int
	primes = append(primes, 2)

	// Build a naive list of integers
	allnums := make([]bool, 1e7)
	for i := range allnums {
		allnums[i] = true
	}
	// Set initial primality for 0, 1, 2
	allnums[0] = false
	allnums[1] = false

	// Quick and dirty implementation of sieve of eratosthenes
	for i := range allnums {
		if allnums[i] {
			primes = append(primes, i)
			for j := i; j*i < 1e7; j++ {
				if i*j < 1e7-1 {
					allnums[i*j] = false
				}
			}
		}
	}

	return primes[n]
}

func main() {
	p := find_nth_prime(10001)
	fmt.Println(p)
}
