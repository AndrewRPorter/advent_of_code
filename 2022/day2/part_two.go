package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Calculate total score if outcome is given.

Outcome of X = lose (0), Y = draw (3) and Z = win (6)

Rock (1 point) => X and A
Paper (2 points) => Y, B
Scissors (3 points) => C, Z
0 points for a loss, 3 for a draw and 6 for a win
*/
func GetScoreForRoundWithOutcome(playerOneMove string, outcome string) int {
	var scoresMap = map[string]int{
		"AX" : 3, // Rock + Loss (0) => Scissors (3)
		"AY" : 4, // Rock + Draw (3) => Rock (1)
		"AZ" : 8, // Rock + Win (6) => Paper (2)
		"BX" : 1, // Paper + Loss (0) => Rock (1)
		"BY" : 5, // Paper + Draw (3) => Paper (2)
		"BZ" : 9, // Paper + Win (6) => Scissors (3)
		"CX" : 2, // Scissors + Loss (0) => Paper (2)
		"CY" : 6, // Scissors + Draw (3) => Scissors (3)
		"CZ" : 7, // Scissors + Win (6) => Rock (1)
	}
	return scoresMap[playerOneMove + outcome]
}

func PartTwo() {
	file, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	totalScore := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")
		playerOneMove, outcome := splits[0], splits[1]
		totalScore += GetScoreForRoundWithOutcome(playerOneMove, outcome)
	}

	fmt.Printf("Total score: %v\n", totalScore)
}
