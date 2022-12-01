package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func getThreeMostCalories(scanner *bufio.Scanner) (int, error) {
	currentElfCalorieTotal := 0
	topThreeCalorieCounts := []int{0, 0, 0}

	for scanner.Scan() {
		/*
			If we are at a new line we should check if the current elf's calorie
			total is greater than the least of the top 3 calorie totals in the
			sorted slice. Note that the sorted slice is in increasing order.
		*/
		if scanner.Text() == "" {
			if currentElfCalorieTotal > topThreeCalorieCounts[0] {
				topThreeCalorieCounts[0] = currentElfCalorieTotal
				// Re-sort the slice in the case that the newly inserted calorie is greater than others
				sort.Ints(topThreeCalorieCounts)
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

	sum := topThreeCalorieCounts[0] + topThreeCalorieCounts[1] + topThreeCalorieCounts[2]
	return sum, nil
}

func PartTwo() {
	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calorieTotal, err := getThreeMostCalories(scanner)
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
