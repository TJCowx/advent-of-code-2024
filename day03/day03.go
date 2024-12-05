package day03

// https://adventofcode.com/2024/day/3

import (
	"advent-of-code-2024/file_reader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var INPUT_PATH = "day03/input.txt"

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

func part1(path string) int64 {
	fmt.Println("DAY 03 PART 1")

	input := file_reader.Read(path)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(input, -1)

	var sum int64 = 0

	for _, match := range matches {
		val1, _ := strconv.ParseInt(match[1], 10, 64)
		val2, _ := strconv.ParseInt(match[2], 10, 64)
		eval := val1 * val2

		sum += eval
	}

	fmt.Printf("RESULT: %d\n", sum)

	return sum
}

func part2(path string) int64 {
	fmt.Println("DAY 03 PART 2")

	input := strings.ToLower(file_reader.Read(path))

	re := regexp.MustCompile(`(?:do(?:n't|)\(\)|$)`)
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var sum int64 = 0
	shouldMul := true
	for len(input) > 0 {
		index := re.FindStringIndex(input)
		eval := input[:index[0]]
		input = input[index[0]:]
		if shouldMul {
			matches := reMul.FindAllStringSubmatch(eval, -1)
			for _, match := range matches {
				val1, _ := strconv.ParseInt(match[1], 10, 64)
				val2, _ := strconv.ParseInt(match[2], 10, 64)
				sum += (val1 * val2)
			}
		}

		shouldMul = strings.HasPrefix(input, "do()")
		if len(input) > 0 {
			input = input[1:]
		}

	}

	fmt.Printf("RESULT: %d\n", sum)

	return sum
}
