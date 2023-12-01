package main

import (
	"fmt"
	"strings"
)

const INPUT_FILE_NAME = "input.txt"

func PrintTopOfStack(mapping map[int]string) {
	for i := 1; i < 10; i++ {
		fmt.Printf("%c", mapping[i][len(mapping[i])-1])
	}
	fmt.Println()
}

func BuildMapping(mappingText []string) map[int]string {
	numStacks := len(strings.Replace(mappingText[len(mappingText)-1], " ", "", -1))

	crateMapping := make(map[int]string)

	// We use len(mappingText)-1 as we want to exclude the last line which
	// shows how many stacks there are (e.g. " 1   2   3   4").
	for crateIndex := 0; crateIndex < len(mappingText)-1; crateIndex++ {
		line := mappingText[crateIndex]
		for stackIndex := 0; stackIndex < numStacks; stackIndex++ {
			if fmt.Sprintf("%c", line[1+stackIndex*4]) == " " {
				continue
			}
			// Create the mapping with a 1 index so we can process moves easily
			crateMapping[stackIndex+1] = fmt.Sprintf("%c", line[1+stackIndex*4]) + crateMapping[stackIndex+1]
		}
	}

	return crateMapping
}

func main() {
	PartOne()
	PartTwo()
}
