package day2

import (
	"fmt"
	"nated/aoc/util"
	"strings"
)

func part1() {
	fmt.Printf("\n------------ PART 1 ------------\n")

	scanner := util.OpenFile("days/day2/part1.txt")
	defer scanner.Close()

	choiceScores := make(map[string]int)
	choiceScores["X"] = 1 // Rock
	choiceScores["Y"] = 2 // Paper
	choiceScores["Z"] = 3 // Scissors

	playMatch := func(win string, draw string, loss string, move string) int {
		if move == win {
			return 6
		} else if move == draw {
			return 3
		} else {
			return 0
		}
	}

	score := 0

	for scanner.Scan() {
		match := strings.Split(scanner.Text(), " ")
		score += choiceScores[match[1]]
		switch match[0] {
		case "A":
			score += playMatch("Y", "X", "Z", match[1])
		case "B":
			score += playMatch("Z", "Y", "X", match[1])
		case "C":
			score += playMatch("X", "Z", "Y", match[1])
		}
	}

	fmt.Printf("Total Score: %d\n", score)
}

func part2() {
	fmt.Printf("\n------------ PART 1 ------------\n")

	scanner := util.OpenFile("days/day2/part2.txt")
	defer scanner.Close()

	resultScores := make(map[string]int)
	resultScores["X"] = 0 // Loss
	resultScores["Y"] = 3 // Draw
	resultScores["Z"] = 6 // Win

	choiceScores := make(map[string]int)
	choiceScores["A"] = 1 // Rock
	choiceScores["B"] = 2 // Paper
	choiceScores["C"] = 3 // Scissors

	playMatch := func(win string, draw string, loss string, move string) int {
		if move == "X" {
			return choiceScores[loss]
		} else if move == "Y" {
			return choiceScores[draw]
		} else {
			return choiceScores[win]
		}
	}

	score := 0

	for scanner.Scan() {
		match := strings.Split(scanner.Text(), " ")
		score += resultScores[match[1]]
		switch match[0] {
		case "A":
			score += playMatch("B", "A", "C", match[1])
		case "B":
			score += playMatch("C", "B", "A", match[1])
		case "C":
			score += playMatch("A", "C", "B", match[1])
		}
	}

	fmt.Printf("Total Score: %d\n", score)
}

func Run() {
	part1()
	part2()
}
