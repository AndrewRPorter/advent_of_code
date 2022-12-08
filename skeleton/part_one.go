package main

import (
	"bufio"
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
	}
}
