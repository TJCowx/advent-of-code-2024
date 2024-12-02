package day02

// https://adventofcode.com/2024/day/2

import (
	"advent-of-code-2024/file_reader"
	"fmt"
	"strconv"
	"strings"
)

var INPUT_PATH = "day02/input.txt"

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

func parse_input(path string) []string {
	content := file_reader.Read(path)

	lines := strings.Split(content, "\n")

	// Get rid of that empty line
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

func part1(path string) int {
	fmt.Println("DAY 02, PART 1")

	lines := parse_input(path)

	var safeLevels int = 0

	for _, line := range lines {
		numStrs := strings.Fields(line)

		asc := false
		desc := false
		safe := true

		lastNum, _ := strconv.Atoi(numStrs[0])

		for i := 1; i < len(numStrs); i++ {
			curr, _ := strconv.Atoi(numStrs[i])
			diff := curr - lastNum

			if !asc && !desc {
				// If we haven't set, lets set
				asc = diff > 0
				desc = !asc
			}

			if asc && (diff <= 0 || diff > 3) {
				safe = false
				break
			} else if desc && (diff >= 0 || diff < -3) {
				safe = false
				break
			}

			lastNum = curr
		}

		if safe {
			safeLevels += 1
		}
	}

	fmt.Printf("PART 1: %d\n", safeLevels)

	return safeLevels
}

func lineStrsToNums(numStrs []string) []int {
	var nums []int

	for _, str := range numStrs {
		converted, _ := strconv.Atoi(str)
		nums = append(nums, converted)
	}

	return nums
}

func allInRange(deltas []int, allowZero bool) bool {
	var posDelts []int
	var negDelts []int
	zeroCount := 0

	for _, delta := range deltas {
		if delta == 0 {
			zeroCount++
			if zeroCount > 1 || !allowZero {
				return false
			}
		} else if delta > 0 && delta <= 3 {
			posDelts = append(posDelts, delta)
		} else if delta < 0 && delta >= -3 {
			negDelts = append(negDelts, delta)
		}
	}

	// Everything is positive or negative and in range, there is only 1 0
	if len(posDelts) == len(deltas) || len(negDelts) == len(deltas) {
		return true
	}

	// Allow for 1 0 count
	if zeroCount == 1 && allowZero && (len(posDelts) == len(deltas)-1 || len(negDelts) == len(deltas)-1) {
		return true
	}

	return false
}

func isSafe(nums []int, allowZeros bool) bool {
	var deltas []int

	for i := 1; i < len(nums); i++ {
		delta := nums[i] - nums[i-1]
		deltas = append(deltas, delta)
	}

	if allInRange(deltas, allowZeros) {
		return true
	}

	return false
}

func part2(path string) int {
	fmt.Println("DAY 02, PART 2")

	lines := parse_input(path)

	var safeLevels int = 0

	for _, line := range lines {
		numStrs := strings.Fields(line)
		nums := lineStrsToNums(numStrs)

		if isSafe(nums, true) {
			safeLevels++
		} else {
			fmt.Println(nums)
			for i := 0; i < len(nums); i++ {
				copyNums := append([]int{}, nums[:i]...)
				copyNums = append(copyNums, nums[i+1:]...)
				fmt.Println(copyNums)
				if isSafe(copyNums, false) {
					safeLevels++
					break
				}
			}

			fmt.Printf("FAILED: %s\n", line)
		}
	}

	fmt.Printf("DAY 2 PART 2: %d\n", safeLevels)

	return safeLevels
}
