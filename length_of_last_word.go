package main

import "fmt"

func main() {
	a := "Hello World"
	fmt.Println(lengthOfLastWord(a))

	a = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(a))

	a = "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(a))
}

func lengthOfLastWord(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i > 0 && s[i] != ' ' && s[i-1] == ' ' {
			result++
			break
		}

		if s[i] == ' ' {
			continue
		}

		result++
		if s[i] == ' ' {
			fmt.Println("skip")
			break;
		}
	}

	return result
}
