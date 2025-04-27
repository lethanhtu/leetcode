package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(3))
	fmt.Println(mySqrt(5))
}

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	start := 1
	end := x
	result := 0

	for start <= end {
		mid := (end + start) / 2

		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			result = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return result
}
