package day05

import "fmt"

var INPUT_PATH = "day05/input.txt"

func Run(part *string) {
	if part == nil {
		part1(INPUT_PATH)
		part2(INPUT_PATH)
	} else if *part == "1" {
		part1(INPUT_PATH)
	} else if *part == "2" {
		part2(INPUT_PATH)
	} else {
		fmt.Println("INVALID INPUT")
	}
}

func part1(path string) int {
	fmt.Println("DAY 05 PART 1")
	fmt.Println("NOT IMPLEMENTED")

	// Map out rules, each number has a set of rules that must be adhered, only if both are in it
	// Data structure???????
	// Map out order

	return 0
}

func part2(path string) int {
	fmt.Println("DAY 05 PART 2")
	fmt.Println("NOT IMPLEMENTED")

	return 0
}
