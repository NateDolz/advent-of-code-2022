package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"nated/aoc/util"
	"strconv"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func part1() {
	fmt.Printf("\n------------ PART 1 ------------\n")
	scanner := util.OpenFile("days/day1/part1.txt")
	defer scanner.Close()

	elf := 0
	calories := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if elf > calories {
				calories = elf
			}
			elf = 0
		} else {
			num, _ := strconv.Atoi(line)
			elf += num
		}
	}

	if elf > calories {
		calories = elf
	}

	fmt.Printf("Most Calories: %d\n", calories)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part2() {
	fmt.Printf("\n------------ PART 2 ------------\n")

	scanner := util.OpenFile("days/day1/part2.txt")
	defer scanner.Close()

	elf := 0
	calories := make([]int, 3)

	checkCalorieCount := func() {
		if idx, min := util.Min(calories); elf > min {
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

func Run() {
	part1()
	part2()
}
