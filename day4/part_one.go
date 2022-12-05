package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func DoSectionsFullyOverlap(sectionOne string, sectionTwo string) bool {
	sectionOneSplits := strings.Split(sectionOne, "-")
	sectionTwoSplits := strings.Split(sectionTwo, "-")

	sectionOneMin, _ := strconv.Atoi(sectionOneSplits[0])
	sectionOneMax, _ := strconv.Atoi(sectionOneSplits[1])

	sectionTwoMin, _ := strconv.Atoi(sectionTwoSplits[0])
	sectionTwoMax, _ := strconv.Atoi(sectionTwoSplits[1])

	sectionOneContainsSectionTwo := sectionOneMin <= sectionTwoMin && sectionOneMax >= sectionTwoMax
	sectionTwoContainsSectionOne := sectionTwoMin <= sectionOneMin && sectionTwoMax >= sectionOneMax

	return sectionOneContainsSectionTwo || sectionTwoContainsSectionOne
}

func PartOne() {
	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0
	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, ",")
		sectionOne, sectionTwo := splits[0], splits[1]

		if DoSectionsFullyOverlap(sectionOne, sectionTwo) {
			counter++
		}
	}

	fmt.Printf("Total overlapping sections: %v\n", counter)
}
