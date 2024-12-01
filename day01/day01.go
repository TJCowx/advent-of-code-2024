package day01

import (
	"advent-of-code-2024/file_reader"
	"fmt"
)

func Run(part *string) {
	if part == nil {
		part1()
		part2()
	} else if *part == "1" {
		part1()
	} else if *part == "2" {
		part2()
	} else {
		fmt.Println("INVALID INPUT")
	}
}

func parse_input() {
	content := file_reader.Read("day01/test-input.txt")

	fmt.Printf(content)
}

func part1() {
	fmt.Println("Day 01, Part 1")
	parse_input()
}

func part2() {
	fmt.Println("Day 01, Part 2")
}
