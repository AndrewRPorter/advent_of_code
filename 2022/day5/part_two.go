package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ProcessMovePartTwo(mapping map[int]string, move string) {
	moveSplits := strings.Split(move, " ")
	numCratesToMove, _ := strconv.Atoi(moveSplits[1])
	fromStack, _ := strconv.Atoi(moveSplits[3])
	toStack, _ := strconv.Atoi(moveSplits[5])

	if numCratesToMove > 1 {
		mapping[toStack] += mapping[fromStack][len(mapping[fromStack])-numCratesToMove:]
		mapping[fromStack] = mapping[fromStack][:len(mapping[fromStack])-numCratesToMove]
	} else {
		mapping[toStack] += fmt.Sprintf("%c", mapping[fromStack][len(mapping[fromStack])-1])
		mapping[fromStack] = mapping[fromStack][:len(mapping[fromStack])-1]
	}
}

func PartTwo() {
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
				ProcessMovePartTwo(mapping, text)
			} else {
				mappingText = append(mappingText, text)
			}
		}
	}

	PrintTopOfStack(mapping)
}
