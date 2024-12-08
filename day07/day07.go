package day07

// https://adventofcode.com/2024/day/7

import (
	"advent-of-code-2024/file_reader"
	"advent-of-code-2024/utils"
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
	Concat
)

func (o *Operand) Fmt() string {
	switch *o {
	case Add:
		return "Add"
	case Multiply:
		return "Multiply"
	case Concat:
		return "Concat"
	default:
		return "I DO NOT KNOW"

	}
}

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

func (e *Evaluation) Solve(isPartOne bool, brute bool) int {
	if isPartOne && brute {
		if evalP1Brute(e.total, 0, e.nums, Add) || evalP1Brute(e.total, 0, e.nums, Multiply) {
			return e.total
		}
	} else if isPartOne && !brute {
		if evalP1Opt(e.total, e.nums, Multiply) || evalP1Opt(e.total, e.nums, Add) {
			return e.total
		}
	} else if brute {
		if evalP2Brute(e.total, 0, e.nums, Add) || evalP2Brute(e.total, 0, e.nums, Multiply) || evalP2Brute(e.total, 0, e.nums, Concat) {
			return e.total
		}

	} else {
		if evalP2Opt(e.total, e.nums, Concat) || evalP2Opt(e.total, e.nums, Multiply) || evalP2Opt(e.total, e.nums, Add) {
			return e.total
		}
	}

	return 0
}

func evalP1Brute(target int, current int, nums []int, operand Operand) bool {
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

	return evalP1Brute(target, currCopy, rest, Add) || evalP1Brute(target, currCopy, rest, Multiply)
}

func evalP1Opt(current int, nums []int, op Operand) bool {
	if len(nums) == 0 {
		return current == 0
	}

	if current <= 0 {
		return false
	}

	num := nums[len(nums)-1]

	currCopy := current
	if op == Add {
		currCopy -= num
	} else {
		// The current number has to be a multiple, otherwise it's not a valid answer
		if current%num != 0 {
			return false
		}
		currCopy /= num
	}

	rest := nums[:len(nums)-1]

	return evalP1Opt(currCopy, rest, Multiply) || evalP1Opt(currCopy, rest, Add)
}

func evalP2Brute(target int, current int, nums []int, operand Operand) bool {
	if len(nums) == 0 {
		return current == target
	}

	if current > target {
		return false
	}

	currCopy := current
	next := nums[0]
	if operand == Add {
		currCopy += next
	} else if operand == Multiply {
		currCopy *= next
	} else {
		concatStr := strconv.Itoa(currCopy) + strconv.Itoa(next)
		concatNum, _ := strconv.Atoi(concatStr)
		currCopy = concatNum
	}

	rest := nums[1:]

	return evalP2Brute(target, currCopy, rest, Add) ||
		evalP2Brute(target, currCopy, rest, Multiply) ||
		evalP2Brute(target, currCopy, rest, Concat)
}

func evalP2Opt(current int, nums []int, op Operand) bool {
	if len(nums) == 0 {
		return current == 0
	}

	if current <= 0 && op != Concat {
		return false
	}

	currCopy := current
	next := nums[len(nums)-1]
	rest := nums[:len(nums)-1]
	if op == Add {
		currCopy -= next
	} else if op == Multiply {
		// The current number has to be a multiple, otherwise it's not a valid answer
		if current%next != 0 {
			return false
		}
		currCopy /= next
	} else {
		currStr := strconv.Itoa(currCopy)
		nextStr := strconv.Itoa(next)
		// We need to end with the number
		if !strings.HasSuffix(currStr, nextStr) {
			return false
		}

		currStr = currStr[:len(currStr)-len(nextStr)]
		currCopy, _ = strconv.Atoi(currStr)
	}

	return evalP2Opt(currCopy, rest, Concat) ||
		evalP2Opt(currCopy, rest, Multiply) ||
		evalP2Opt(currCopy, rest, Add)
}

func part1(path string) (int, int) {
	fmt.Println("DAY 07 PART 1")
	content := file_reader.Read(path)
	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sum := 0

	timer := utils.BuildTimer()
	timer.Start()
	for _, line := range lines {
		eval := buildEval(line)
		sum += eval.Solve(true, true)
	}
	timer.End()

	fmt.Println("BRUTE FORCE")
	fmt.Printf("RESULT: %d | RUN TIME: %s\n", sum, timer.TimeLapsed())

	sumOpt := 0

	timer.Start()
	for _, line := range lines {
		eval := buildEval(line)
		sumOpt += eval.Solve(true, false)
	}
	timer.End()

	fmt.Println("OPTIMIZED")
	fmt.Printf("RESULT: %d | RUN TIME: %s\n", sumOpt, timer.TimeLapsed())

	return sum, sumOpt
}

func part2(path string) (int, int) {
	fmt.Println("DAY 07 PART 2")

	content := file_reader.Read(path)
	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	sum := 0

	timer := utils.BuildTimer()
	timer.Start()
	for _, line := range lines {
		eval := buildEval(line)
		sum += eval.Solve(false, true)
	}
	timer.End()

	fmt.Println("BRUTE FORCE")
	fmt.Printf("RESULT: %d | RUN TIME: %s\n", sum, timer.TimeLapsed())

	sumOpt := 0

	timer.Start()
	for _, line := range lines {
		eval := buildEval(line)
		sumOpt += eval.Solve(false, false)
	}
	timer.End()

	fmt.Println("OPTIMIZED")
	fmt.Printf("RESULT: %d | RUN TIME: %s\n", sumOpt, timer.TimeLapsed())

	return sum, sumOpt
}
