package main

import "fmt"

func main() {
	var str []string
	str = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(str))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	firstStr := strs[0]
	result := ""
outerLoop:
	for i, _ := range firstStr {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 {
				break outerLoop
			}

			if firstStr[i] != strs[j][i] {
				break outerLoop
			}

		}

		result = result + string(firstStr[i])

	}
	return result
}
