package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"nate-aoc/util"
	"strconv"

	"golang.org/x/exp/slices"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func part1() {
	fmt.Printf("\n------------ PART 1 ------------\n")
	scanner := util.OpenFile("part1.txt")
	defer scanner.Close()

	elf := 0
	calories := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && elf > calories {
			calories = elf
			elf = 0
		} else if line != "" {
			num, _ := strconv.Atoi(line)
			elf += num
		} else {
			elf = 0
		}
	}

	fmt.Printf("Most Calories: %d\n", calories)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part2() {
	fmt.Printf("\n------------ PART 2 ------------\n")
	scanner := util.OpenFile("part2.txt")
	defer scanner.Close()

	elf := 0
	calories := make([]int, 3)

	checkCalorieCount := func() {
		if idx := slices.IndexFunc(calories, func(calorie int) bool { return calorie == 0 }); idx != -1 {
			calories[idx] = elf
		} else if idx, min := util.Min(calories); elf > min {
			calories[idx] = elf
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			checkCalorieCount()
			elf = 0
		} else {
			num, _ := strconv.Atoi(line)
			elf += num
		}
	}

	if elf > 0 {
		checkCalorieCount()
	}

	fmt.Printf("Top 3 Calorie Total: %d\n", calories[0]+calories[1]+calories[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	part1()
	part2()
}
