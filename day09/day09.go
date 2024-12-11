package day09

// https://adventofcode.com/2024/day/9

import (
	"advent-of-code-2024/file_reader"
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

var INPUT_PATH = "day09/input.txt"

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

func isEven(num int) bool {
	return num%2 == 0
}

func makeArr(val int, length int) []int {
	newArr := make([]int, length)

	for i := range newArr {
		newArr[i] = val
	}

	return newArr
}

func buildInput(path string) []int {
	input := strings.Split(file_reader.ReadIntoStrArr(path)[0], "")
	var mappedOutput []int
	fileNum := 0

	for i, char := range input {
		num, _ := strconv.Atoi(char)

		if isEven(i) {
			mappedOutput = append(mappedOutput, makeArr(fileNum, num)...)
			fileNum += 1
		} else {
			blanks := make([]int, num)

			for j := 0; j < num; j++ {
				blanks[j] = -1
			}

			mappedOutput = append(mappedOutput, blanks...)
		}
	}

	return mappedOutput
}

func isEmpty(val int) bool {
	return val == -1
}

func isNotEmpty(val int) bool {
	return val != -1
}

func solveBruteP1(path string) int {
	input := buildInput(path)

	sum := 0

	for i := 0; i < len(input); i++ {
		curr := input[i]
		if curr != -1 {
			sum += curr * i
			continue
		}

		lastFilledI, hasFilled := utils.FindLastIndex(input, isNotEmpty)

		if lastFilledI <= i || !hasFilled {
			break
		}

		next := input[lastFilledI]
		sum += next * i
		input[lastFilledI] = -1
	}

	return sum
}

func lastSame(slice []int, target int, start int) int {
	for i := start; i >= 0; i-- {
		if slice[i] != target {
			return i
		}
	}

	return -1
}

func nextFilledI(slice []int, start int) int {
	for i := start; i < len(slice); i++ {
		if slice[i] != -1 {
			return i
		}
	}

	return -1
}

func findAvailable(slice []int, requiredSpace int) (int, bool) {
	count := 0

	for i, val := range slice {
		if val == -1 {
			count++
			if count == requiredSpace {
				return i - requiredSpace, true
			}
		} else {
			count = 0
		}
	}

	return -1, false
}

func solveBruteP2(path string) int {
	// Note, the best way to do this is in building the input
	// but i don't care here weeeeeeeeeeeeeee
	input := buildInput(path)

	sum := 0

	for i := len(input) - 1; i >= 0; {
		if input[i] == -1 {
			i--
			continue
		}

		val := input[i]
		// Find the length we need here
		lastSameI := lastSame(input, val, i-1)

		requiredSpace := i - lastSameI

		// Find the first group of available space that allows this to fit
		startAvailI, hasAvail := findAvailable(input, requiredSpace)

		if !hasAvail || startAvailI >= i {
			i -= requiredSpace

			continue
		}

		// Insert the values into that space and do the math
		for j := startAvailI + 1; j <= startAvailI+requiredSpace; j++ {
			input[j] = val
		}

		// Remove the values
		for j := i; j > i-requiredSpace; j-- {
			input[j] = -1
		}

		// Set the index to the next last space
		i -= requiredSpace
	}

	// Could I do math in the above loop? Yes. Will I? No
	for i, val := range input {
		if val == -1 {
			continue
		}

		sum += val * i
	}

	return sum
}

func part1(path string) int {
	fmt.Println("DAY 09 PART 1")

	timer := utils.BuildTimer()
	timer.Start()

	res := solveBruteP1(path)

	timer.End()

	fmt.Printf("RESULT %d | TIME ELAPSED: %s\n", res, timer.TimeLapsed())

	return res
}

func part2(path string) int {
	fmt.Println("DAY 09 PART 2")

	timer := utils.BuildTimer()
	timer.Start()

	res := solveBruteP2(path)

	timer.End()

	fmt.Printf("RESULT %d | TIME ELAPSED: %s\n", res, timer.TimeLapsed())

	return res
}
