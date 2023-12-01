package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func CalculateGroupPriority(rucksackOne string, rucksackTwo string, rucksackThree string) (int, error) {
	for _, char  := range(rucksackOne) {
		// Find first intersection char and calculate the priority based
		// on the unicode value and offset from alphabet.
		stringChar := fmt.Sprintf("%c", char)
		if strings.Contains(rucksackTwo, stringChar) && strings.Contains(rucksackThree, stringChar) {
			if unicode.IsLower(char) {
				// lowercase unicode alphabet goes from 97-122
				return int(char) - 96, nil
			} else {
				// uppercase unicode alphabet goes from 65-90
				// Note that uppercase chars have priority of
				// 27-53 so we add 26
				return (int(char) - 64) + 26, nil
			}
		}
	}

	err := errors.New("Unable to find an intersection")
	
	return 0, err
}

func PartTwo() {
	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	totalScore := 0
	counter := 0
	rucksacks := []string{"", "", ""}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if counter == 3 {
			score, err := CalculateGroupPriority(rucksacks[0], rucksacks[1], rucksacks[2])
			if err != nil {
				log.Fatal(err)
				return
			}

			totalScore += score
			counter = 0
		}
		rucksacks[counter] = scanner.Text()
		counter += 1
	}

	// I know this is gross...
	// Calculate one more time with the remaining items in the rucksack slice
	score, err := CalculateGroupPriority(rucksacks[0], rucksacks[1], rucksacks[2])
	if err != nil {
		log.Fatal(err)
		return
	}

	totalScore += score

	fmt.Printf("Total score: %v\n", totalScore)
}
