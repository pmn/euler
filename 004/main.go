/*
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func reverse(str string) string {
	var out string
	for i := len(str) - 1; i >= 0; i-- {
		out += string(str[i])
	}
	return out
}

func is_palindromic(num int) bool {
	str := strconv.Itoa(num)
	return str == reverse(str)
}

func main() {
	biggest := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if i*j > biggest && is_palindromic(i*j) {
				biggest = i * j
			}
		}
	}

	fmt.Println(biggest)
}
