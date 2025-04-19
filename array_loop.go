package main

import (
	"fmt"
)

func main() {
	var input = [6]int{7, 10, 11, 14, 26}
	sum := 24

	input[5] = 10
	fmt.Println(input)
	fmt.Println(findSolution(input[:], sum))

}

func findSolution(nums []int, target int) []int {
	waitingMap := map[int]int{}

	for i, v := range nums {
		for waitingIndex, waitingValue := range waitingMap {
			if waitingValue == v {
				return []int{i, waitingIndex}
			}
		}

		waitingMap[i] = target - v
	}

	return []int{}
}
