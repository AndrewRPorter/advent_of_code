package main

import "fmt"

const INPUT_FILE_NAME = "input.txt"

func stringInSlice(a string, slice []string) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}

func IsSequenceOfUniqueChars(text string, n int) bool {
	slice := []string{}
	for i := 0; i < n; i++ {
		if !stringInSlice(fmt.Sprintf("%c", text[i]), slice) {
			slice = append(slice, fmt.Sprintf("%c", text[i]))
		} else {
			return false
		}
	}
	return true
}

func main() {
	PartOne()
	PartTwo()
}
