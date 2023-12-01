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

func CalculatePriorityOfIntersection(compartmentOne string, compartmentTwo string) (int, error) {
	for _, char  := range(compartmentOne) {
		// Find first intersection char and calculate the priority based
		// on the unicode value and offset from alphabet.
		if strings.Contains(compartmentTwo, fmt.Sprintf("%c", char)) {
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

func PartOne() {
	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	totalScore := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// grab rucksack from input text file and split into
		// it's two evenly sized compartments
		rucksackInput := scanner.Text()
		compartmentOne := rucksackInput[0:len(rucksackInput)/2]
		compartmentTwo := rucksackInput[len(rucksackInput)/2:]

		score, err := CalculatePriorityOfIntersection(compartmentOne, compartmentTwo)
		if err != nil {
			log.Fatal(err)
			return
		}

		totalScore += score
	}

	fmt.Printf("Total score: %v\n", totalScore)
}
