package day01

// https://adventofcode.com/2024/day/1

import (
	"advent-of-code-2024/file_reader"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var INPUT_PATH = "day01/input.txt"

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

func parse_lists(path string) ([]int, []int) {
	content := file_reader.Read(path)

	lines := strings.Split(content, "\n")

	// Get rid of that empty line
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var listA []int
	var listB []int

	for _, line := range lines {
		inputs := strings.Fields(line)
		a, b := inputs[0], inputs[1]
		numA, _ := strconv.Atoi(a)
		numB, _ := strconv.Atoi(b)
		listA = append(listA, numA)
		listB = append(listB, numB)
	}

	return listA, listB
}

func part1(path string) int {
	fmt.Println("Day 01, Part 1")

	listA, listB := parse_lists(path)

	sort.Ints(listA)
	sort.Ints(listB)

	var sum int = 0

	for i := 0; i < len(listA); i++ {
		diff := listA[i] - listB[i]

		sum += int(math.Abs(float64(diff)))
	}

	fmt.Printf("ANSWER IS: %d\n", sum)
	return sum
}

func part2(path string) int {
	fmt.Println("Day 01, Part 2")

	listA, listB := parse_lists(path)

	countMap := map[int]int{}

	for _, item := range listB {
		val, exists := countMap[item]
		if exists {
			countMap[item] = val + 1
		} else {
			countMap[item] = 1
		}
	}

	var sum int = 0

	for _, item := range listA {
		count, exists := countMap[item]
		if exists {
			sum += (item * count)
		}
	}

	fmt.Printf("ANSWER IS %d\n", sum)

	return sum
}
