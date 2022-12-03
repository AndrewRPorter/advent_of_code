package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Calculate player 2's score for a game of rock papers scissors.

Rock (1 point) => X and A
Paper (2 points) => Y, B
Scissors (3 points) => C, Z
0 points for a loss, 3 for a draw and 6 for a win
*/
func GetScoreForRound(playerOneMove string, playerTwoMove string) int {
	var scoresMap = map[string]int{
		"AX" : 4, // Draw (3) + Rock (1)
		"AY" : 8, // Win (6) + Paper (2)
		"AZ" : 3, // Loss (0) + Scissors (3)
		"BX" : 1, // Loss (0) + Rock (1)
		"BY" : 5, // Draw (3) + Paper (2)
		"BZ" : 9, // Win (6) + Scissors (3)
		"CX" : 7, // Win (6) + Rock (1)
		"CY" : 2, // Loss (0) + Paper (2)
		"CZ" : 6, // Draw (3) + Scissors (3)
	}
	return scoresMap[playerOneMove + playerTwoMove]
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
		line := scanner.Text()
		splits := strings.Split(line, " ")
		playerOneMove, playerTwoMove := splits[0], splits[1]
		totalScore += GetScoreForRound(playerOneMove, playerTwoMove)
	}

	fmt.Printf("Total score: %v\n", totalScore)
}
