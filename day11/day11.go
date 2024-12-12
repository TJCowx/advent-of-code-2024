package day11

// https://adventofcode.com/2024/day/11

import (
	"advent-of-code-2024/file_reader"
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

var INPUT_PATH = "day11/input.txt"

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

func parseInput(path string) []int {
	input := file_reader.Read(path)

	var vals []int

	for _, val := range strings.Split(strings.TrimSpace(input), " ") {
		num, _ := strconv.Atoi(val)
		vals = append(vals, num)
	}

	return vals
}

func splitStone(stoneNum int) []int {
	asStr := strconv.Itoa(stoneNum)

	mid := len(asStr) / 2

	num1, _ := strconv.Atoi(asStr[:mid])
	num2, _ := strconv.Atoi(asStr[mid:])

	return []int{num1, num2}
}

func blinkStone(stoneNum int) []int {
	// Just split the stone into 2 if it's even
	if stoneNum == 0 {
		return []int{1}
	}

	// If we have an even num of digits, split it in half
	if len(strconv.Itoa(stoneNum))%2 == 0 {
		return splitStone(stoneNum)
	}

	if stoneNum >= 9223372036854775807/2024 {
		fmt.Printf("GOING TO BE OVERFLOW %d", stoneNum)
	}

	return []int{stoneNum * 2024}
}

func part1(path string) int {
	fmt.Println("DAY 11 PART 1")
	stones := parseInput(path)

	timer := utils.BuildTimer()

	timer.Start()
	for i := 0; i < 25; i++ {
		var blinkedStones []int

		for _, stone := range stones {
			blinkedStones = append(blinkedStones, blinkStone(stone)...)
		}

		stones = blinkedStones

		fmt.Printf("RUN %d\n", i+1)
	}

	res := len(stones)

	timer.End()

	fmt.Printf("RESULT: %d | TIME ELAPSED: %s\n", res, timer.TimeLapsed())

	return res
}

func part2(path string) int {
	fmt.Println("DAY 11 PART 2")

	return 0
}
