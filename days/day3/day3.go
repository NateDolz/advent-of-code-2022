package day3

import (
	"fmt"
	"log"
	"nated/aoc/util"
	"strings"
)

var giftValues = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func part1() {
	fmt.Printf("\n------------ PART 1 ------------\n")

	value := 0
	scanner := util.OpenFile("days/day3/part1.txt")
	for scanner.Scan() {
		line := []rune(scanner.Text())
		idx := len(line) / 2
		part1 := line[:idx]
		part2 := line[idx:]
		union := util.Union(part1, part2)

		if len(union) > 1 {
			fmt.Println(line)
			log.Fatalln("Too many matches on line")
		}

		value += strings.Index(giftValues, string(union[0])) + 1
	}

	fmt.Printf("Total gift value: %d\n", value)
}

func Run() {
	part1()
}
