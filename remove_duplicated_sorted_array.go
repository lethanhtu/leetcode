package main

import "fmt"

func main() {
	nums := []int{}

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))

	nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {
	i := 0
	for i < len(nums)-1 {
		j := i + 1
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}

		if j == len(nums) {
			break
		}

		for k := i + 1; k < j; k++ {
			nums[k] = nums[j]
		}

		i = i + 1
	}

	for i = 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			break
		}
	}

	return i + 1
}
