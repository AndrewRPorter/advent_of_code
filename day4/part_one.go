package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PartOne() {
	fmt.Println("part 1")

	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}
