package main

import "fmt"

func main() {
	var x int = 12321
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	var digits []int

	for x > 0 {
		digits = append(digits, x%10)
		x = x / 10
	}

	size := len(digits)
	for i := 0; i < size; i++ {
		if digits[i] != digits[size-i-1] {
			return false
		}
	}

	return true
}
