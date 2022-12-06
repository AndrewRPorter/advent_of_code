package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ProcessMovePartOne(mapping map[int]string, move string) {
	moveSplits := strings.Split(move, " ")
	numCratesToMove, _ := strconv.Atoi(moveSplits[1])
	fromStack, _ := strconv.Atoi(moveSplits[3])
	toStack, _ := strconv.Atoi(moveSplits[5])

	for i := 0; i < numCratesToMove; i++ {
		mapping[toStack] += fmt.Sprintf("%c", mapping[fromStack][len(mapping[fromStack])-1])
		mapping[fromStack] = mapping[fromStack][:len(mapping[fromStack])-1]
	}
}

func PartOne() {
	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var mapping map[int]string

	mappingText := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			mapping = BuildMapping(mappingText)
		} else {
			if strings.Contains(text, "move") {
				ProcessMovePartOne(mapping, text)
			} else {
				mappingText = append(mappingText, text)
			}
		}
	}

	PrintTopOfStack(mapping)
}
