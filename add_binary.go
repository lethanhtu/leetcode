package main

import "fmt"

func main() {
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	indexA := len(a) - 1
	indexB := len(b) - 1
	result := ""
	modifier := "0"
	cA := "0"
	cB := "0"
	current := "0"

	for indexA >= 0 || indexB >= 0 {
		fmt.Println(indexA, indexB)
		if indexA >= 0 {
			cA = string(a[indexA])
			indexA = indexA - 1
		}

		fmt.Println("after-A")

		if indexB >= 0 {
			cB = string(b[indexB])
			indexB = indexB - 1
		}

		fmt.Println("after-B")

		if cA == "1" && cB == "1" {
			if modifier == "1" {
				current = "1"
				modifier = "1"
			} else {
				current = "0"
				modifier = "1"
			}
		} else if cA == "1" || cB == "1" {
			if modifier == "1" {
				current = "0"
			} else {
				current = "1"
			}
		} else {
			if modifier == "1" {
				current = "1"
				modifier = "0"
			} else {
				current = "0"
			}
		}

		result = current + result
		cA = "0"
		cB = "0"
	}

	if modifier == "1" {
		result = "1" + result
	}

	return result
}
