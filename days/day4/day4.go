package day4

import (
	"fmt"
	"nated/aoc/util"
	"strconv"
	"strings"
)

func isEnveloping(start1 int, end1 int, start2 int, end2 int) bool {
	// 2 contains 1
	if start1 >= start2 && end1 <= end2 {
		return true
	} else if start2 >= start1 && end2 <= end1 {
		return true
	}
	return false
}

func isOverlapping(start1 int, end1 int, start2 int, end2 int) bool {
	if enveloping := isEnveloping(start1, end1, start2, end2); enveloping {
		return true
	} else if start1 <= start2 && start2 <= end1 {
		return true
	} else if start2 <= start1 && start1 <= end2 {
		return true
	} else if start1 >= start2 && start2 >= end1 {
		return true
	} else if start2 >= start1 && start1 >= end2 {
		return true
	}
	return false
}

func getAssignmentRanges(assignments []string) (start1 int, end1 int, start2 int, end2 int) {
	assignment1 := strings.Split(assignments[0], "-")
	assignment2 := strings.Split(assignments[1], "-")

	x1, _ := strconv.Atoi(assignment1[0])
	y1, _ := strconv.Atoi(assignment1[1])
	x2, _ := strconv.Atoi(assignment2[0])
	y2, _ := strconv.Atoi(assignment2[1])

	return x1, y1, x2, y2
}

func part1() {
	fmt.Printf("\n------------ PART 1 ------------\n")

	value := 0
	scanner := util.OpenFile("days/day4/part1.txt")
	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		start1, end1, start2, end2 := getAssignmentRanges(assignments)

		if envelops := isEnveloping(start1, end1, start2, end2); envelops {
			value++
		}
	}

	fmt.Printf("Total overlaps: %d\n", value)
}

func part2() {
	fmt.Printf("\n------------ PART 2 ------------\n")

	value := 0
	scanner := util.OpenFile("days/day4/part2.txt")
	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		start1, end1, start2, end2 := getAssignmentRanges(assignments)

		if envelops := isOverlapping(start1, end1, start2, end2); envelops {
			value++
		}
	}

	fmt.Printf("Total overlaps: %d\n", value)
}

func Run() {
	part1()
	part2()
}
