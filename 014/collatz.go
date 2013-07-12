// The following iterative sequence is defined for the set of positive integers:
//
// n -> n/2 (n is even)
// n -> 3n + 1 (n is odd)
//
// Using the rule above and starting with 13, we generate the following sequence:
//
// 13  40  20  10  5  16  8  4  2  1
//
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
// Although it has not been proved yet (Collatz Problem), it is thought that all starting
// numbers finish at 1.
//
// Which starting number, under one million, produces the longest chain?
//
// NOTE: Once the chain starts the terms are allowed to go above one million.

package main

import "fmt"

type chain struct {
	start  int
	length int
}

func calc(number int) int {
	if number%2 == 0 {
		return number / 2
	}
	return 3*number + 1
}

func collatz(start int) chain {
	c := chain{start, 1}
	i := int(start)

	for i > 1 {
		c.length++
		i = calc(i)
	}
	return c
}

func find_longest() {
	longest := chain{0, 0}

	for i := 0; i < 1e6; i++ {
		current := collatz(i)
		if current.length > longest.length {
			longest.start = current.start
			longest.length = current.length
		}
	}

	fmt.Println(longest)
}

func main() {
	find_longest()
}
