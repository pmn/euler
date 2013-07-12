/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

func main() {
	found := false
	counter := 20
	for found == false {
		success := true
		for i := 2; i <= 20; i++ {
			if counter%i != 0 {
				counter++
				success = false
			}
		}
		found = success
	}

	fmt.Println(counter)
}
