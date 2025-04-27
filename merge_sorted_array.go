package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	//fmt.Println(nums1)
}

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	k := m + n - 1
// 	i := m - 1
// 	j := n - 1

// 	for i >= 0 {
// 		if j >= 0 && nums2[j] > nums1[i] {
// 			nums1[k] = nums2[j]
// 			j--
// 		} else {
// 			nums1[k] = nums1[i]
// 			i--
// 		}
// 		k--
// 	}
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	i := m - 1
	j := n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
		fmt.Println(nums1)
	}
}
