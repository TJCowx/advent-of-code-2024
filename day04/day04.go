package day04

// https://adventofcode.com/2024/day/4

import (
	"advent-of-code-2024/file_reader"
	"advent-of-code-2024/utils"
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

func checkNeighbourP1(board [][]string, x int, y int, dir Direction, checkingChar string) (int, [][]int) {
	nextX, nextY := nextIndex(x, y, dir)

	if !isInBounds(board, nextX, nextY) {
		return 0, [][]int{}
	}

	if board[nextY][nextX] == checkingChar {
		if checkingChar == "S" {
			return 1, [][]int{{nextX, nextY}, {x, y}}
		}
		val, coords := checkNeighbourP1(board, nextX, nextY, dir, nextChar(checkingChar))

		if val == 1 {
			coords = append(coords, []int{x, y})
			return val, coords
		}
	}

	return 0, [][]int{}
}

func checkNeighbourForX(board [][]string, x int, y int) (int, [][]int) {
	// Check if top-left has S or M
	topLeftX, topLeftY := nextIndex(x, y, UpLeft)
	if !isInBounds(board, topLeftX, topLeftY) {
		return 0, [][]int{}
	}
	tl := board[topLeftY][topLeftX]
	if tl != "S" && tl != "M" {
		return 0, [][]int{}
	}
	// Check if bottom-right has the other
	bottomRightX, bottomRightY := nextIndex(x, y, DownRight)
	if !isInBounds(board, bottomRightX, bottomRightY) {
		return 0, [][]int{}
	}
	br := board[bottomRightY][bottomRightX]
	if (br == "S" && tl != "M") || (br == "M" && tl != "S") || (br != "M" && br != "S") {
		return 0, [][]int{}
	}

	// Check if top-right as S or M
	topRightX, topRightY := nextIndex(x, y, UpRight)
	if !isInBounds(board, topRightX, topRightY) {
		return 0, [][]int{}
	}
	tr := board[topRightY][topRightX]
	if tr != "S" && tr != "M" {
		return 0, [][]int{}
	}
	// Check if bototm-left has the other
	bottomLeftX, bottomLeftY := nextIndex(x, y, DownLeft)
	if !isInBounds(board, bottomLeftX, bottomLeftY) {
		return 0, [][]int{}
	}
	bl := board[bottomLeftY][bottomLeftX]

	if (bl == "S" && tr != "M") || (bl == "M" && tr != "S") || (bl != "M" && bl != "S") {
		return 0, [][]int{}
	}

	return 1, [][]int{
		{topLeftX, topLeftY},
		{x, y},
		{bottomRightX, bottomRightY},
		{topRightX, topRightY},
		{bottomLeftX, bottomLeftY},
	}
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

func printCoords(board [][]string, coords [][]int) {
	if len(coords) == 0 {
		return
	}
	var str = ""

	for _, coord := range coords {
		str += board[coord[1]][coord[0]]
	}

	fmt.Printf("MATCH: %s\n", str)
}

func part1(path string) int {
	fmt.Println("DAY 04 PART 1")
	input := file_reader.Read(path)

	timer := utils.BuildTimer()
	timer.Start()

	board := parseInput(input)
	sum := 0

	// DEBUGGING
	// pPrint(board)
	allSafeCoords := make(map[string]bool)

	for y, row := range board {
		for x, item := range row {
			if item == "X" {
				for _, dir := range DIRECTIONS {
					val, safeCoords := checkNeighbourP1(board, x, y, dir, "M")
					allSafeCoords = addToMap(allSafeCoords, safeCoords)
					sum += val
				}
			}
		}
	}

	// DEBUGGING
	// removeNonSafe(board, allSafeCoords)

	timer.End()
	fmt.Printf("RESULT: %d | TIME ELAPSED: %s\n", sum, timer.TimeLapsed())

	return sum
}

func part2(path string) int {
	fmt.Println("DAY 04 PART 2")
	input := file_reader.Read(path)

	timer := utils.BuildTimer()
	timer.Start()

	board := parseInput(input)
	sum := 0

	// DEBUGGING
	//pPrint(board)
	allSafeCoords := make(map[string]bool)

	for y, row := range board {
		for x, item := range row {
			if item == "A" {
				// Check top left
				val, safeCoords := checkNeighbourForX(board, x, y)
				allSafeCoords = addToMap(allSafeCoords, safeCoords)

				sum += val
			}
		}
	}

	// DEBUGGING
	// removeNonSafe(board, allSafeCoords)

	timer.End()
	fmt.Printf("RESULT: %d | TIME ELAPSED: %s\n", sum, timer.TimeLapsed())

	return sum
}
