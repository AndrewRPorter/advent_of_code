package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PartOne() {
	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		for i, _ := range text {
			if IsSequenceOfUniqueChars(text[i:], 4) {
				fmt.Printf("Start of sequence at: %v\n", i+4)
				break
			} else {
				continue
			}
		}
	}
}
