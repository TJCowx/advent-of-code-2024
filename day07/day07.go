package day07

import (
	"advent-of-code-2024/file_reader"
	"fmt"
	"strconv"
	"strings"
)

var INPUT_PATH = "day07/input.txt"

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

type Evaluation struct {
	total int
	nums  []int
}

type Operand int

const (
	Add Operand = iota
	Multiply
)

func buildEval(line string) Evaluation {
	sects := strings.Split(line, ":")

	// First half
	total, _ := strconv.Atoi(sects[0])
	var nums []int

	trimmed := strings.TrimSpace(sects[1])
	numStrs := strings.Split(trimmed, " ")

	for _, str := range numStrs {
		num, _ := strconv.Atoi(str)

		nums = append(nums, num)
	}

	return Evaluation{total, nums}
}

func (e *Evaluation) Solve() int {
	if evaluate(e.total, 0, e.nums, Add) || evaluate(e.total, 0, e.nums, Multiply) {
		return e.total
	}

	return 0
}

func evaluate(target int, current int, nums []int, operand Operand) bool {
	if len(nums) == 0 {
		return current == target
	}

	currCopy := current
	if operand == Add {
		currCopy += nums[0]
	} else {
		currCopy *= nums[0]
	}

	if current > target {
		return false
	}

	rest := nums[1:]

	return evaluate(target, currCopy, rest, Add) || evaluate(target, currCopy, rest, Multiply)
}

func part1(path string) int {
	fmt.Println("DAY 07 PART 1")
	content := file_reader.Read(path)
	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sum := 0

	for _, line := range lines {
		eval := buildEval(line)
		sum += eval.Solve()
	}

	fmt.Printf("RESULT: %d\n", sum)

	return sum
}

func part2(path string) int {
	fmt.Println("DAY 07 PART 2")

	return 0
}
