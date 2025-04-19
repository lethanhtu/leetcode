package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(16))
}

func mySqrt(x int) int {
	if x == 1 || x == 2 {
		return 1
	}

	for i := 1; i < x/2; i++ {
		if i*i == x {
			return i
		}

		if i*i > x {
			return i - 1
		}
	}

	return 0
}
