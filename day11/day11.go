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

func splitStone(stoneNum int) (int, int) {
	asStr := strconv.Itoa(stoneNum)

	mid := len(asStr) / 2

	num1, _ := strconv.Atoi(asStr[:mid])
	num2, _ := strconv.Atoi(asStr[mid:])

	return num1, num2
}

func buildKey(stone int, remainingBlinks int) string {
	return strconv.Itoa(stone) + "-" + strconv.Itoa(remainingBlinks)
}

func solve(stone int, remainingBlinks int, cache map[string]int) int {
	if remainingBlinks == 0 {
		return 1
	}

	key := buildKey(stone, remainingBlinks)

	if val, exists := cache[key]; exists {
		return val
	}

	sum := 0

	if stone == 0 {
		sum += solve(1, remainingBlinks-1, cache)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		subStone1, subStone2 := splitStone(stone)

		sum += solve(subStone1, remainingBlinks-1, cache)
		sum += solve(subStone2, remainingBlinks-1, cache)
	} else {
		sum += solve(stone*2024, remainingBlinks-1, cache)
	}

	cache[key] = sum
	return sum
}

func part1(path string) int {
	fmt.Println("DAY 11 PART 1")
	stones := parseInput(path)

	timer := utils.BuildTimer()

	timer.Start()

	stoneCache := make(map[string]int)
	sum := 0
	for _, stone := range stones {
		sum += solve(stone, 25, stoneCache)
	}

	timer.End()

	fmt.Printf("RESULT: %d | TIME ELAPSED: %s\n", sum, timer.TimeLapsed())

	return sum
}

func part2(path string) int {
	fmt.Println("DAY 11 PART 2")
	stones := parseInput(path)

	timer := utils.BuildTimer()

	timer.Start()

	stoneCache := make(map[string]int)
	sum := 0
	for _, stone := range stones {
		sum += solve(stone, 75, stoneCache)
	}

	timer.End()

	fmt.Printf("RESULT: %d | TIME ELAPSED: %s\n", sum, timer.TimeLapsed())

	return sum
}
