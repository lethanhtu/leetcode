package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	result := 0

	// nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	// val = 2
	// result = 0
	// fmt.Println(nums)
	// result = removeElement(nums, val)
	// fmt.Println(result)
	// fmt.Println(nums)

	// nums = []int{3, 2, 2, 3}
	// val = 2
	// fmt.Println(nums)
	// result = removeElement(nums, val)
	// fmt.Println(result)
	// fmt.Println(nums)

	nums = []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 10}
	val = 2
	fmt.Println(nums)
	result = removeElement(nums, val)
	fmt.Println(result)
	fmt.Println(nums)

	// nums = []int{2, 2, 2, 2, 2, 2, 2, 2}
	// val = 2
	// fmt.Println(nums)
	// result = removeElement(nums, val)
	// fmt.Println(result)
	// fmt.Println(nums)

	// nums = []int{2}
	// val = 2
	// fmt.Println(nums)
	// result = removeElement(nums, val)
	// fmt.Println(result)
	// fmt.Println(nums)

	// nums = []int{2, 2}
	// val = 2
	// fmt.Println(nums)
	// result = removeElement(nums, val)
	// fmt.Println(result)
	// fmt.Println(nums)

	// nums = []int{3, 3}
	// val = 5
	// fmt.Println(nums)
	// result = removeElement(nums, val)
	// fmt.Println(result)
	// fmt.Println(nums)

	// nums = []int{4, 5}
	// val = 5
	// fmt.Println(nums)
	// result = removeElement(nums, val)
	// fmt.Println(result)
	// fmt.Println(nums)

}

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}

		fmt.Println(nums)
	}

	return k
}

func __removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize two pointers

	// Continue until pointers meet
	left := 0
	right := len(nums) - 1

	// Continue until pointers meet

	for left < right {
		// If right pointer points to target, decrement it
		if nums[right] == val {
			right--
			continue
		}

		// If left pointer points to target, swap with right pointer
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}

		// Move left pointer forward
		left++
	}

	count := 0

	for count < len(nums) && nums[count] != val {
		count++
	}

	return count

}

func _removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	start := 0
	end := len(nums) - 1
	count := len(nums) - 1

	for start < end {
		for nums[end] == val {
			end--
			if end < 0 {
				break
			}
		}

		for nums[start] != val {
			start++
			if start == len(nums) {
				break
			}
		}

		if start == len(nums) || end < 0 || start > end {
			break
		}

		nums[start] = nums[end]
		nums[end] = val

		end--
		start++
	}

	for nums[count] == val {
		count = count - 1
		if count < 0 {
			break
		}
	}

	return count + 1
}
