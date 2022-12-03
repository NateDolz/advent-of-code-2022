package main

import (
	"fmt"
	"log"
	"nated/aoc/days/day1"
	"nated/aoc/days/day2"
	"os"
	"strings"
)

func runProgram(program string) {
	fmt.Printf("\n!=!=!=!=!=!=!=!=!= %s !=!=!=!=!=!=!=!=!=\n", program)
	switch strings.ToLower(program) {
	case "day1":
		day1.Run()
	case "day2":
		day2.Run()
	default:
		log.Printf("Could not run program %s\n", program)
	}
}

func main() {
	programs := os.Args[1:]
	for _, prog := range programs {
		runProgram(prog)
	}
}