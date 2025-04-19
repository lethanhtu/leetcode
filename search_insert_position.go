package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5

	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 2

	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target && (mid == len(nums)-1 || nums[mid+1] > target) {
			return mid + 1
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return -1
}
