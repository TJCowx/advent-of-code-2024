package day06

// https://adventofcode.com/2024/day/6

import (
	"advent-of-code-2024/file_reader"
	"fmt"
	"strconv"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Left
	Right
	Down
)

var INPUT_PATH = "day06/input.txt"

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

type Board struct {
	tiles        [][]string
	width        int
	height       int
	blockedTiles map[string]bool
	X            int
	Y            int
	dir          Direction
	Visited      map[string]bool
}

func getBoard(path string) Board {
	content := file_reader.Read(path)
	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	startX := 0
	startY := 0
	var tiles [][]string
	width := len([]rune(lines[0])) - 1
	height := len(lines) - 1
	blockedTiles := make(map[string]bool)

	for y, line := range lines {
		runes := []rune(line)
		for x, r := range runes {
			if r == '#' {
				blockedTiles[getKey(x, y)] = true
			} else if r == '^' {
				startY = y
				startX = x
			}

		}
	}

	return Board{
		tiles:        tiles,
		width:        width,
		height:       height,
		blockedTiles: blockedTiles,
		X:            startX,
		Y:            startY,
		dir:          Up,
		Visited:      make(map[string]bool),
	}
}

func (b *Board) ChangeDir() {
	switch b.dir {
	case Up:
		b.dir = Right
	case Right:
		b.dir = Down
	case Down:
		b.dir = Left
	default:
		b.dir = Up
	}
}

func getKey(x int, y int) string {
	return strconv.Itoa(x) + "-" + strconv.Itoa(y)
}

func (b *Board) Print() {
	fmt.Println("===============BOARD================")
	for _, row := range b.tiles {
		fmt.Println(row)
	}
}

func (b *Board) NextGuardSpot() (int, int) {
	switch b.dir {
	case Up:
		return b.X, b.Y - 1
	case Right:
		return b.X + 1, b.Y
	case Down:
		return b.X, b.Y + 1
	default:
		return b.X - 1, b.Y
	}
}

func (b *Board) MoveGuard(x int, y int) {
	b.X = x
	b.Y = y
}

func (b *Board) HasGuardEscaped() bool {
	return b.X > b.width || b.X < 0 || b.Y > b.height || b.Y < 0
}

func (b *Board) IsNextSafe(x int, y int) bool {
	if b.X > b.width || b.X < 0 || b.Y > b.height || b.Y < 0 {
		return true
	}

	if b.blockedTiles[getKey(x, y)] {
		return false
	}

	return true
}

func part1(path string) int {
	fmt.Println("DAY 06 PART 1")
	content := file_reader.Read(path)
	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	board := getBoard(path)

	fmt.Printf("(%d, %d)", board.X, board.Y)
	board.Print()

	for {
		board.Visited[getKey(board.X, board.Y)] = true

		nextX, nextY := board.NextGuardSpot()

		// Turn until next is safe
		for !board.IsNextSafe(nextX, nextY) {
			board.ChangeDir()
			nextX, nextY = board.NextGuardSpot()
		}

		board.MoveGuard(nextX, nextY)

		if board.HasGuardEscaped() {
			break
		}
	}

	res := len(board.Visited)

	fmt.Printf("RESULT: %d\n", res)

	return res
}

func part2(path string) int {
	fmt.Println("DAY 06 PART 1")
	fmt.Println("NOT IMPLEMENTED")

	return 0
}
