package day04

import (
	"advent-of-code-2024/file_reader"
	"fmt"
	"strings"
)

var INPUT_PATH string = "day04/input.txt"

type Direction int

const (
	UpLeft Direction = iota
	Up
	UpRight
	Left
	Right
	DownLeft
	Down
	DownRight
)

var DIRECTIONS = []Direction{
	UpLeft, Up, UpRight, Left, Right, DownLeft, Down, DownRight,
}

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

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")

	// Get rid of that empty line
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var rows [][]string

	for _, line := range lines {
		var row []string
		runes := []rune(line)

		for _, rune := range runes {
			row = append(row, string(rune))
		}

		rows = append(rows, row)
	}

	return rows
}

func nextChar(curr string) string {
	switch curr {
	case "X":
		return "M"
	case "M":
		return "A"
	case "A":
		return "S"
	default:
		fmt.Printf("FOUND OTHER: %s", curr)
		return ""
	}
}

func prevChar(curr string) string {
	switch curr {
	case "S":
		return "A"
	case "A":
		return "M"
	case "M":
		return "X"
	default:
		fmt.Printf("FOUND OTHER: %s", curr)
		return ""
	}
}

func nextIndex(currX int, currY int, dir Direction) (int, int) {
	x := currX
	y := currY

	switch dir {
	case UpLeft:
		x -= 1
		y -= 1
	case Up:
		y -= 1
	case UpRight:
		x += 1
		y -= 1
	case Left:
		x -= 1
	case Right:
		x += 1
	case DownLeft:
		x -= 1
		y += 1
	case Down:
		y += 1
	case DownRight:
		x += 1
		y += 1
	}

	return x, y
}

func isInBounds(board [][]string, x int, y int) bool {
	if y >= len(board) || y < 0 {
		return false
	}

	if x >= len(board[y]) || x < 0 {
		return false
	}

	return true
}

func checkNeighbour(board [][]string, x int, y int, dir Direction, checkingChar string) (int, [][]int) {
	nextX, nextY := nextIndex(x, y, dir)

	if !isInBounds(board, nextX, nextY) {
		return 0, [][]int{}
	}

	if board[nextY][nextX] == checkingChar {
		if checkingChar == "S" {
			return 1, [][]int{{x, y}, {nextX, nextY}}
		}
		val, coords := checkNeighbour(board, nextX, nextY, dir, nextChar(checkingChar))

		if val == 1 {
			coords = append(coords, []int{x, y})
			return val, coords
		}
	}

	return 0, [][]int{}
}

func pPrint(board [][]string) {
	fmt.Println("====================")
	for _, row := range board {
		fmt.Printf("[ %v ]\n", row)
	}
}

func addToMap(inMap map[string]bool, vals [][]int) map[string]bool {
	outMap := inMap

	for _, coords := range vals {
		key := fmt.Sprintf("%d,%d", coords[0], coords[1])
		outMap[key] = true
	}

	return outMap
}

func removeNonSafe(board [][]string, safeCoords map[string]bool) {
	boardCopy := board

	for y, row := range board {
		for x := range row {
			coords := fmt.Sprintf("%d,%d", x, y)
			if !safeCoords[coords] {
				boardCopy[y][x] = "."
			}
		}
	}

	pPrint(boardCopy)
}

func part1(path string) int {
	fmt.Println("DAY 04 PART 1")
	input := file_reader.Read(path)
	board := parseInput(input)
	sum := 0

	// DEBUGGING
	// pPrint(board)
	allSafeCoords := make(map[string]bool)

	for y, row := range board {
		for x, item := range row {
			if item == "X" {
				for _, dir := range DIRECTIONS {
					val, safeCoords := checkNeighbour(board, x, y, dir, "M")
					allSafeCoords = addToMap(allSafeCoords, safeCoords)
					sum += val
				}
			}
		}
	}

	// DEBUGGING
	// removeNonSafe(board, allSafeCoords)

	fmt.Printf("RESULT: %d\n", sum)

	return sum
}

func part2(path string) int {
	fmt.Println("DAY 04 PART 2")
	fmt.Println("NOT IMPLEMENTED")

	return 0
}
