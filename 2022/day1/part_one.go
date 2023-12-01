package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getMostCalories(scanner *bufio.Scanner) (int, error) {
	currentElfCalorieTotal := 0
	maxElfCalorieTotal := 0

	for scanner.Scan() {
		/*
			If we are at a new line we should check if the current elf's calorie
			total is greater than the current maximum. If so, update the max calorie count.
			We also want to reset the current elf's calorie total
		*/
		if scanner.Text() == "" {
			if currentElfCalorieTotal > maxElfCalorieTotal {
				maxElfCalorieTotal = currentElfCalorieTotal
			}
			currentElfCalorieTotal = 0
		} else {
			lineString := scanner.Text()
			calories, err := strconv.Atoi(lineString)
			if err != nil {
				return 0, err
			}
			currentElfCalorieTotal += calories
		}
	}

	return maxElfCalorieTotal, nil
}

func PartOne() {
	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calorieTotal, err := getMostCalories(scanner)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Most calories: %v\n", calorieTotal)
}
