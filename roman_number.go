package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	mapping := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for i := 0; i < len(s); i++ {
		//if i > len(s) {
		//	break
		//}

		if i == len(s)-1 {
			result = result + mapping[s[i]]
			return result
		}

		if mapping[s[i]] >= mapping[s[i+1]] {

			result = result + mapping[s[i]]
			continue
		}

		result = result + mapping[s[i+1]] - mapping[s[i]]
		i = i + 1
	}

	return result
}
