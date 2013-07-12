package main

import "testing"

func TestIsPalindromic(t *testing.T) {
	if is_palindromic(666) == false {
		t.Fail()
	}
}

func TestIsNotPalindromic(t *testing.T) {
	if is_palindromic(665) {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	test := "test"
	if reverse(test) != "tset" {
		t.Fail()
	}
}
