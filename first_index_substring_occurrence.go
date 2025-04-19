package main

import "fmt"

func main() {
	fmt.Println(strStr("abcdef123bc", "123"))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}
