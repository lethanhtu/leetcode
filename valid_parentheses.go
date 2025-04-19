package main

import "fmt"

func main() {
	str := "[{}]"
	fmt.Println(isValid(str))
}

func isValid(s string) bool {
	stack := make([]rune, 0)

	closeToOpen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			lastElement := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if closeToOpen[c] != lastElement {
				return false
			}
		}
	}

	return len(stack) == 0
}
