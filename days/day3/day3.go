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
	defer scanner.Close()

	for scanner.Scan() {
		line := []rune(scanner.Text())
		idx := len(line) / 2
		part1 := line[:idx]
		part2 := line[idx:]
		intersection := util.Intersect(part1, part2)

		if len(intersection) > 1 {
			fmt.Println(line)
			log.Fatalln("Too many matches on line")
		}

		value += strings.Index(giftValues, string(intersection[0])) + 1
	}

	fmt.Printf("Total gift value: %d\n", value)
}

func part2() {
	fmt.Printf("\n------------ PART 2 ------------\n")

	value := 0
	idx := 0
	groups := make([][]rune, 2)

	scanner := util.OpenFile("days/day3/part2.txt")
	defer scanner.Close()

	for scanner.Scan() {
		line := []rune(scanner.Text())
		if idx == 2 {
			intersection := util.Intersect(groups[0], groups[1])
			intersection = util.Intersect(intersection, line)
			if len(intersection) > 1 {
				fmt.Println(line)
				log.Fatalln("Too many matches on line")
			}
			value += strings.Index(giftValues, string(intersection[0])) + 1
			idx = 0
		} else {
			groups[idx] = line
			idx++
		}
	}

	fmt.Printf("Total gift value: %d\n", value)
}

func Run() {
	part1()
	part2()
}
