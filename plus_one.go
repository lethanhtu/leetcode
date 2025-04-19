package main

import "fmt"

func main() {
	a := []int{1,2,3,4}
	fmt.Println(plusOne(a))

	a = []int{1,2,3,9}
	fmt.Println(plusOne(a))

	a = []int{1,1,1,1,1,1,1,9}
	fmt.Println(plusOne(a))

	a = []int{9, 0, 9}
	fmt.Println(plusOne(a))

	a = []int{9}
	fmt.Println(plusOne(a))

	a = []int{9,9}
	fmt.Println(plusOne(a))

	a = []int{9,9,9}
	fmt.Println(plusOne(a))

}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	return append([]int{1}, digits...)
}